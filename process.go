package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const osHaveSigTerm = true

func ShellInvocationCommand(interactive bool, root, command string) []string {
	shellArgument := "-c"
	if interactive {
		shellArgument = "-ic"
	}
	profile := filepath.Join(root, ".profile")
	shellCommand := fmt.Sprintf("source \"%s\" 2>/dev/null; %s", profile, command)
	return []string{"/bin/bash", shellArgument, shellCommand}

}

type Process struct {
	Command     string
	Interactive bool

	*exec.Cmd
}

func NewProcess(workdir, command string, interactive bool) (p *Process) {
	argv := ShellInvocationCommand(interactive, workdir, command)
	return &Process{
		command, interactive, exec.Command(argv[0], argv[1:]...),
	}
}

func (p *Process) Start() error {
	p.PlatformSpecificInit()
	return p.Cmd.Start()
}

func (p *Process) Signal(signal syscall.Signal) error {
	group, err := os.FindProcess(-1 * p.Process.Pid)
	if err == nil {
		err = group.Signal(signal)
	}
	return err
}

func (p *Process) PlatformSpecificInit() {
	if !p.Interactive {
		p.SysProcAttr = &syscall.SysProcAttr{}
		p.SysProcAttr.Setsid = true
	}
	return
}

func (p *Process) SendSigTerm() {
	p.Signal(syscall.SIGTERM)
}

func (p *Process) SendSigKill() {
	p.Signal(syscall.SIGKILL)
}
