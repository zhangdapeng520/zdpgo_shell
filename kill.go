package zdpgo_shell

import (
	"fmt"
	"runtime"
)

//Kill 终止指定进程
func (s *Shell) Kill(pid int) string {
	var result string
	if runtime.GOOS == "windows" {
		result = s.Run(fmt.Sprintf("taskkill /pid %d  -t  -f", pid))
	} else {
		result = s.Run(fmt.Sprintf("kill -9 %d", pid))
	}
	return result
}
