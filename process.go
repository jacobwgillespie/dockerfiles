package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const osHaveSigTerm = true

func ShellInvocationCommand(root, command string) []string {
	shellArgument := "-c"
	profile := filepath.Join(root, ".profile")
	shellCommand := fmt.Sprintf("source \"%s\" 2>/dev/null; %s", profile, command)
	return []string{"/bin/bash", shellArgument, shellCommand}

}

type Process struct {
	Command string

	*exec.Cmd
}

func NewProcess(workdir, command string) (p *Process) {
	argv := ShellInvocationCommand(workdir, command)
	return &Process{
		command, exec.Command(argv[0], argv[1:]...),
	}
}

func (p *Process) Start() error {
	p.SysProcAttr = &syscall.SysProcAttr{}
	p.SysProcAttr.Setsid = true
	return p.Cmd.Start()
}

func (p *Process) Signal(signal syscall.Signal) error {
	group, err := os.FindProcess(-1 * p.Process.Pid)
	if err == nil {
		err = group.Signal(signal)
	}
	return err
}

func (p *Process) SendSigTerm() {
	p.Signal(syscall.SIGTERM)
}

func (p *Process) SendSigKill() {
	p.Signal(syscall.SIGKILL)
}
