package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
)

type Manager struct {
	outletFactory *OutletFactory

	teardown, teardownNow Barrier // signal shutting down

	wg sync.WaitGroup
}

func (m *Manager) monitorInterrupt() {
	handler := make(chan os.Signal, 1)
	signal.Notify(handler, os.Interrupt)

	first := true

	for sig := range handler {
		switch sig {
		case os.Interrupt:
			fmt.Println("      | ctrl-c detected")

			m.teardown.Fall()
			if !first {
				m.teardownNow.Fall()
			}
			first = false
		}
	}
}

func (m *Manager) startProcess(idx int, name string, args []string, of *OutletFactory) {
	const interactive = false
	workDir := filepath.Dir(".")
	ps := NewProcess(workDir, strings.Join(args, " "))

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
	go of.LineReader(pipeWait, name, idx, stdout, false)
	go of.LineReader(pipeWait, name, idx, stderr, true)

	finished := make(chan struct{}) // closed on process exit

	err = ps.Start()
	if err != nil {
		m.teardown.Fall()
		of.SystemOutput(fmt.Sprint("Failed to start ", name, ": ", err))
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
		case <-finished:
			return

		case <-m.teardown.Barrier():
			// Manager tearing down

			if !osHaveSigTerm {
				of.SystemOutput(fmt.Sprintf("Killing %s", name))
				ps.Process.Kill()
				return
			}

			of.SystemOutput(fmt.Sprintf("sending SIGTERM to %s", name))
			ps.SendSigTerm()

			// Give the process a chance to exit, otherwise kill it.
			select {
			case <-m.teardownNow.Barrier():
				of.SystemOutput(fmt.Sprintf("Killing %s", name))
				ps.SendSigKill()
			case <-finished:
			}
		}
	}()
}
