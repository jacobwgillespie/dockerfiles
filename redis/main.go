package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/go-etcd/etcd"
)

const shutdownGraceTime = 20 * time.Second

const sentinelCfgTmpl = `port %d
sentinel announce-ip %s
sentinel monitor master %s %d 2
sentinel down-after-milliseconds master 60000
sentinel failover-timeout master 180000
sentinel parallel-syncs master 1
`

var (
	redisPort   int
	redisConfig string

	sentinelPort   int
	sentinelConfig string

	publishHost string
	publishPort int

	etcdHost string
	etcdTTL  int

	e *etcd.Client
)

func main() {
	// Parse flags
	flag.IntVar(&redisPort, "redis-port", 6379, "port redis is running on")
	flag.StringVar(&redisConfig, "redis-config", "/etc/redis/redis.conf", "path to redis config")

	flag.IntVar(&sentinelPort, "sentinel-port", 26379, "port sentinel is running on")
	flag.StringVar(&sentinelConfig, "sentinel-config", "/etc/redis/sentinel.conf", "path to sentinel config")

	flag.StringVar(&publishHost, "publish-host", "127.0.0.1", "the host to publish on")
	flag.IntVar(&publishPort, "publish-port", 6379, "the port to publish on")

	flag.StringVar(&etcdHost, "etcd-host", "http://localhost:4001", "the hostname of the etcd node")
	flag.IntVar(&etcdTTL, "etcd-ttl", 10, "etcd ttl in seconds")

	flag.Parse()

	// Set up manager
	of := NewOutletFactory()
	of.Padding = 14

	// Configure etcd
	of.SystemOutput("Connecting to etcd at " + etcdHost)
	e = etcd.NewClient([]string{etcdHost})
	_, err := e.CreateDir("/redis", 0)
	if err != nil {
		of.ErrorOutput(err.Error())
		if strings.Contains(err.Error(), "All the given peers are not reachable") {
			os.Exit(1)
		}
	}
	_, err = e.CreateDir("/redis/cluster", 0)
	if err != nil {
		of.ErrorOutput(err.Error())
	}
	_, err = e.CreateDir("/redis/cluster/nodes", 0)
	if err != nil {
		of.ErrorOutput(err.Error())
	}

	// Read hostname
	hostname, err := os.Hostname()
	handleError(err)

	// Find or elect master
	master := ""
	r, err := e.Get("/redis/cluster/master", false, false)
	if err != nil {
		of.ErrorOutput(err.Error())
	}
	if err != nil || r.Node == nil {
		of.SystemOutput("There is no master - holding an election...")
		_, err := e.Create("/redis/cluster/election", hostname, uint64(etcdTTL))
		if err != nil {
			of.SystemOutput("I lost the election - waiting for the master")
			for {
				r, err := e.Get("/redis/cluster/master", false, false)
				if err == nil && r.Node != nil && r.Node.Value != "" {
					master = r.Node.Value
					of.SystemOutput("Master ready!")
					break
				}
				of.SystemOutput("Waiting for the master...")
				time.Sleep(1 * time.Second)
			}
		} else {
			of.SystemOutput("I won the election!")
		}
	} else {
		master = r.Node.Value
		of.SystemOutput("There is already a cluster master")
	}

	m := &Manager{
		outletFactory: of,
	}

	go m.monitorInterrupt()

	m.teardown.FallHook = func() {
		go func() {
			time.Sleep(shutdownGraceTime)
			of.SystemOutput("Grace time expired")
			m.teardownNow.Fall()
		}()
	}

	// Run redis as master or slave
	if master == "" {
		of.SystemOutput("Starting redis in master mode")
		m.startProcess(1, "redis-server", []string{"redis-server", redisConfig}, of)
	} else {
		of.SystemOutput("Starting redis in slave mode")
		m.startProcess(1, "redis-server", []string{"redis-server", redisConfig, "--slaveof", master}, of)
	}

	// Wait for redis to come alive
	err = nil
	var c net.Conn
	host := fmt.Sprintf("%s:%d", publishHost, redisPort)
	for {
		of.SystemOutput("Looking for redis server on " + host)
		c, err = net.Dial("tcp", host)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	c.Close()

	// Run redis sentinel
	of.SystemOutput("Starting redis sentinel")
	if master == "" {
		ioutil.WriteFile(sentinelConfig, []byte(fmt.Sprintf(sentinelCfgTmpl, sentinelPort, publishHost, publishHost, publishPort)), 0644)
	} else {
		parts := strings.Split(master, " ")
		host := parts[0]
		port, _ := strconv.ParseInt(parts[1], 10, 0)
		ioutil.WriteFile(sentinelConfig, []byte(fmt.Sprintf(sentinelCfgTmpl, sentinelPort, publishHost, host, port)), 0644)
	}
	m.startProcess(2, "redis-sentinel", []string{"redis-sentinel", sentinelConfig}, of)

	// Update etcd
	go func() {
		for {
			if master == "" {
				e.Set("/redis/cluster/master", fmt.Sprintf("%s %d", publishHost, publishPort), uint64(etcdTTL))
			}

			if hostname != "" {
				e.Set("/redis/cluster/nodes/"+hostname, fmt.Sprintf("%s %d", publishPort), uint64(etcdTTL))
			}

			time.Sleep(time.Duration(etcdTTL/2) * time.Second)
		}
	}()

	// Wait
	<-m.teardown.Barrier()
	m.wg.Wait()
}
