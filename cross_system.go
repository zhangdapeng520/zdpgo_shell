package zdpgo_shell

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

/*
@Time : 2022/4/30 8:37
@Author : 张大鹏
@File : cross_system
@Software: Goland2021.3.1
@Description: 跨平台兼容问题解决
*/

// Result 命令执行结果
type Result struct {
	Result string `json:"result"`
	Pid    int    `json:"pid"`
	Error  string `json:"error"`
}

func runInLinux(cmd string) Result {
	var result Result

	// 执行命令
	command := exec.Command("/bin/sh", "-c", cmd)

	// 进程pid
	result.Pid = command.Process.Pid

	// 命令输出
	output, err := command.Output()
	if err != nil {
		result.Error = err.Error()
		return result
	}
	result.Result = strings.TrimSpace(string(output))

	// 返回结果
	return result
}

func runInWindows(cmd string) Result {
	var result Result

	// 创建命令
	command := exec.Command("cmd", "/c", cmd)
	if command.Process != nil {
		result.Pid = command.Process.Pid
	}

	// 输出结果
	output, err := command.Output()
	if err != nil {
		result.Error = err.Error()
		return result
	}
	resultStr := strings.TrimSpace(string(output))

	// 编码转换，防止乱码
	resultStr = ConvertByte2String([]byte(resultStr), "GB18030")
	result.Result = resultStr

	// 返回结果
	return result
}

// Run 执行CMD命令
func Run(cmd string) Result {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

// RunRealTime 获取命令实时结果
func RunRealTime(command string) Result {
	var (
		cmd    *exec.Cmd
		result Result
	)

	// 创建cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", command)
	case "linux":
		cmd = exec.Command("/bin/sh", "-c", command)
	default:
		cmd = exec.Command("/bin/sh", "-c", command)
	}
	if cmd.Process != nil {
		result.Pid = cmd.Process.Pid
	}

	// 创建管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		result.Error = err.Error()
		return result
	}

	// 执行命令并获取输出
	cmd.Start()
	var results []string
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
		fmt.Println(cmdRe)
		results = append(results, cmdRe)
	}

	// 等待命令结束
	err = cmd.Wait()
	if err != nil {
		result.Error = err.Error()
		return result
	}

	// 拼接结果
	resultStr := strings.Join(results, "\n")
	result.Result = resultStr

	// 返回
	return result
}
