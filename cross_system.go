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

func runInLinux(cmd string) string {
	fmt.Println("Running Linux cmd:", cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}

func runInWindows(cmd string) string {
	fmt.Println("Running Win cmd:", cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	resultStr := strings.TrimSpace(string(result))
	resultStr = ConvertByte2String([]byte(resultStr), "GB18030")
	return resultStr
}

func RunCommand(cmd string) string {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

func runInLinuxWithErr(cmd string) (string, error) {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func runInWindowsWithErr(cmd string) (string, error) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	resultStr := strings.TrimSpace(string(result))
	resultStr = ConvertByte2String([]byte(resultStr), "GB18030")
	return resultStr, err
}

func RunWithWait(command string) (string, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", command)
	case "linux":
		cmd = exec.Command("/bin/sh", "-c", command)
	default:
		cmd = exec.Command("/bin/sh", "-c", command)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Start()
	var results []string
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
		// TODO: 实时写入到控制台
		fmt.Println(cmdRe)
		results = append(results, cmdRe)
	}
	cmd.Wait()

	if err != nil {
		return "", err
	}
	resultStr := strings.Join(results, "\n")
	return resultStr, err
}

func RunCommandWithErr(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindowsWithErr(cmd)
	} else {
		return runInLinuxWithErr(cmd)
	}
}
