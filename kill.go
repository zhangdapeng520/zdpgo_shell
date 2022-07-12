package zdpgo_shell

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
)

// Kill 终止指定进程
func (s *Shell) Kill(pid int) Result {
	var result Result
	if runtime.GOOS == "windows" {
		result = s.Run(fmt.Sprintf("taskkill /pid %d  -t  -f", pid))
	} else {
		result = s.Run(fmt.Sprintf("kill -9 %d", pid))
	}
	return result
}

// 删除.env中的进程ID
func (s *Shell) KillEnv() Result {
	var result Result

	// 读取进程数据
	data, err := ioutil.ReadFile(".env")
	if err != nil {
		result.Error = err.Error()
		return result
	}

	// 获取进程ID
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		result.Error = err.Error()
		return result
	}

	// 杀死该进程
	if runtime.GOOS == "windows" {
		result = s.Run(fmt.Sprintf("taskkill /pid %d  -t  -f", pid))
	} else {
		result = s.Run(fmt.Sprintf("kill -9 %d", pid))
	}
	return result
}
