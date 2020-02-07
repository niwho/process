package process

import (
	"os/exec"
)

type Process struct {
	cmd  *exec.Cmd
	name string
	args []string
}

func NewProcess(name string, arg ...string) *Process {
	return &Process{
		name: name,
		args: arg,
	}
}

func (proc *Process) GetPid() int {
	if proc.cmd == nil {
		return -1
	}
	return proc.cmd.Process.Pid
}

func (proc *Process) Stop() error {
	if proc.cmd != nil {
		return proc.cmd.Process.Kill()
	}
	return nil
}

func (proc *Process) Restart() error {
	if proc.cmd == nil {
		return nil
	}

	_ = proc.Stop()
	return proc.cmd.Run()
}

func (proc *Process) Run() (out []byte, err error) {
	if proc.cmd != nil {
		proc.Stop()
	}
	proc.cmd = exec.Command(proc.name, proc.args...)
	out, err = proc.cmd.CombinedOutput()
	return
}

func (proc *Process) GetStat() error {
	panic("implement me")
}
