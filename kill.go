package zdpgo_shell

import (
	"fmt"
	"runtime"
)

//Kill 终止指定进程
func (s *Shell) Kill(pid int) {
	if runtime.GOOS == "windows" {
		s.Run(fmt.Sprintf("taskkill /pid %d  -t  -f", pid))
	} else {
		s.Run(fmt.Sprintf("kill -9 %d", pid))
	}
}
