package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
)

type Manager struct {
	outletFactory *OutletFactory

	teardown, teardownNow Barrier

	wg sync.WaitGroup
}

func (m *Manager) monitorInterrupt() {
	handler := make(chan os.Signal, 1)
	signal.Notify(handler, os.Interrupt)

	first := true

	for sig := range handler {
		switch sig {
		case os.Interrupt:
			fmt.Println("ctrl-c detected")

			m.teardown.Fall()
			if !first {
				m.teardownNow.Fall()
			}
			first = false
		}
	}
}

func (m *Manager) startProcess(idx int, name string, args []string) {
	workDir, err := os.Getwd()
	if err != nil {
		handleError(err)
	}

	const interactive = false
	ps := NewProcess(workDir, strings.Join(args, " "), interactive)

	ps.Stdin = nil

	stdout, err := ps.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := ps.StderrPipe()
	if err != nil {
		panic(err)
	}

	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go m.outletFactory.LineReader(pipeWait, name, idx, stdout, false)
	go m.outletFactory.LineReader(pipeWait, name, idx, stderr, true)

	finished := make(chan struct{})

	err = ps.Start()
	if err != nil {
		m.teardown.Fall()
		m.outletFactory.SystemOutput(fmt.Sprint("Failed to start ", name, ": ", err))
		return
	}

	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		defer close(finished)
		pipeWait.Wait()
		ps.Wait()
	}()

	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		// Prevent goroutine from exiting before process has finished.
		defer func() { <-finished }()
		defer m.teardown.Fall()

		select {
		case <-m.teardown.Barrier():
			if !osHaveSigTerm {
				m.outletFactory.SystemOutput(fmt.Sprintf("Killing %s", name))
				ps.Process.Kill()
				return
			}

			m.outletFactory.SystemOutput(fmt.Sprintf("sending SIGTERM to %s", name))
			ps.SendSigTerm()

			// Give the process a chance to exit, otherwise kill it.
			select {
			case <-m.teardownNow.Barrier():
				m.outletFactory.SystemOutput(fmt.Sprintf("Killing %s", name))
				ps.SendSigKill()
			case <-finished:
			}
		}
	}()

}
