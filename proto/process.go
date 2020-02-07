package proto

type IProcess interface {
	GetPid() int
	Stop()error
	Restart()error
	GetStat()error
}

type IProcessManager interface {
	Run(name string, arg ...string) (proc IProcess, err error) //
	Daemon(name string, arg ...string) error               //
	FindProcess(keyword string) IProcess
}
