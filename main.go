package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"

	"github.com/coreos/go-etcd/etcd"
)

const shutdownGraceTime = 20 * time.Second

const sentinelCfgTmpl = `port %i
sentinel announce-ip %s
sentinel monitor local %s %i 2
sentinel down-after-milliseconds local 60000
sentinel failover-timeout local 180000
sentinel parallel-syncs local 1
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

func handleError(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func main() {
	flag.IntVar(&redisPort, "redis-port", 6379, "port redis is running on")
	flag.StringVar(&redisConfig, "redis-config", "/etc/redis/redis-server.conf", "path to redis config")

	flag.IntVar(&sentinelPort, "sentinel-port", 26379, "port sentinel is running on")
	flag.StringVar(&sentinelConfig, "sentinel-config", "/etc/redis/sentinel.conf", "path to sentinel config")

	flag.StringVar(&publishHost, "publish-host", "127.0.0.1", "the host to publish on")
	flag.IntVar(&publishPort, "publish-port", 6370, "the port to publish on")

	flag.StringVar(&etcdHost, "etcd-host", "http://localhost:4001", "the hostname of the etcd node")
	flag.IntVar(&etcdTTL, "etcd-ttl", 10, "etcd ttl in seconds")

	flag.Parse()

	e = etcd.NewClient([]string{etcdHost})

	e.CreateDir("/redis", 0)
	e.CreateDir("/redis/cluster", 0)
	e.CreateDir("/redis/cluster/nodes", 0)

	hostname, err := os.Hostname()
	handleError(err)

	master := ""

	r, err := e.Get("/redis/cluster/master", false, false)
	if err != nil || r.Node == nil {
		_, err := e.Create("/redis/cluster/election", hostname, 20)
		if err != nil {
			// lost election, wait for master
			for {
				r, err := e.Get("/redis/cluster/master", false, false)
				if err == nil && r.Node != nil && r.Node.Value != "" {
					fmt.Print("\n")
					master = r.Node.Value
					break
				}
				fmt.Print(".")
				time.Sleep(1 * time.Second)
			}
		} else {
			// won election
		}
	} else {
		master = r.Node.Value
	}

	run(master)
}

func run(master string) {
	of := NewOutletFactory()
	of.Padding = 14
	m := &Manager{
		outletFactory: of,
	}

	go m.monitorInterrupt()

	m.teardown.FallHook = func() {
		go func() {
			time.Sleep(shutdownGraceTime)
			fmt.Println("Grace time expired")
			m.teardownNow.Fall()
		}()
	}

	if master == "" {
		m.startProcess(1, "redis-server", []string{"redis-server", redisConfig})
	} else {
		m.startProcess(1, "redis-server", []string{"redis-server", redisConfig, "--slaveof", master})
	}

	// wait for redis to come alive
	var err error
	var c net.Conn
	for {
		c, err = net.Dial("tcp", fmt.Sprintf("%s:6379", publishHost))
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	c.Close()

	ioutil.WriteFile(sentinelConfig, []byte(fmt.Sprintf(sentinelCfgTmpl, sentinelPort, publishHost, publishHost, publishPort)), 0644)
	m.startProcess(2, "redis-sentinel", []string{"redis-sentinel", sentinelConfig})

	go func() {
		for {
			if master == "" {
				e.Set("/redis/cluster/master", fmt.Sprintf("%s 6379", publishHost), uint64(etcdTTL))
			}

			hostname, _ := os.Hostname()
			if hostname != "" {
				e.Set("/redis/cluster/nodes/"+hostname, fmt.Sprintf("%s 6379", publishHost), uint64(etcdTTL))
			}

			time.Sleep(time.Duration(etcdTTL/2) * time.Second)
		}
	}()

	<-m.teardown.Barrier()

	m.wg.Wait()
}
