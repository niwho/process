package process

import (
	"fmt"
	"github.com/niwho/process/proto"
	"github.com/shirou/gopsutil/process"
	"strings"
	"time"
)

type ProcessManager struct {
}

func (pm *ProcessManager) Run(name string, arg ...string) (proc proto.IProcess, err error) {
	panic("implement me")
}

func (pm *ProcessManager) Daemon(duplicateKill bool, shouldGoon func(stdout string) bool, name string, arg ...string) (err error) {
	pps, _ := pm.ListProcessWithFilter(func(cmdLine string) bool {
		if !strings.Contains(cmdLine, name) {
			return false
		}

		for _, ar := range arg {
			if !strings.Contains(cmdLine, ar) {
				return false
			}
		}
		return true
	})
	if len(pps) > 0 {
		if !duplicateKill {

			return nil // 是否杀掉重启
		} else {
			for _, ps := range pps {
				_ = ps.Kill()
			}
		}
	}
	for ; ; {

		proc := NewProcess(name, arg...)
		out, err := proc.Run()
		if err != nil {
			fmt.Println("err", err)
			//return err
		}
		fmt.Println("out: ", string(out))
		//if strings.Contains(string(out), "broken pipe") {
		if shouldGoon(string(out)) {
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
	return nil
}

func (pm *ProcessManager) FindProcess(keyword string) proto.IProcess {
	panic("implement me")
}

func (pm *ProcessManager) ListProcess() proto.IProcess {
	prs, err := process.Processes()
	fmt.Println("err", err)
	for _, p := range prs {
		cmdLine, _ := p.Cmdline()
		name, _ := p.Name()
		fmt.Println(p.Pid, name, cmdLine)
	}
	return nil
}

func (pm *ProcessManager) ListProcessWithFilter(matched func(cmdLine string) bool) (procs []*process.Process, err error) {
	prs, err := process.Processes()
	if err != nil {
		return
	}
	for _, p := range prs {
		cmdLine, err := p.Cmdline()
		if err != nil {
			continue
		}
		if matched(cmdLine) {
			procs = append(procs, p)
		}
	}
	return
}
