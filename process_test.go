package process

import (
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestA1(t *testing.T) {
	p := &ProcessManager{}
	p.ListProcess()
}


func TestDaemon(t *testing.T) {
	p := &ProcessManager{}
	_ = p.Daemon(false, func(stdout string) bool {
		if strings.Contains(stdout, "out0018"){
			return true
		}
		return false
	},"ffmpeg", "-i" ,"out0018.ts", "-vf", "\"transpose=1\"", "ttt2.ts")
}