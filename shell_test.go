package zdpgo_shell

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/4/29 21:58
@Author : 张大鹏
@File : shell_test
@Software: Goland2021.3.1
@Description: 测试核心功能
*/

func getShell() *Shell {
	return New()
}

func TestShell_ExecShell(t *testing.T) {
	s := getShell()
	tables := []string{
		"dir",
		"pwd",
	}
	for _, cmd := range tables {
		result := s.Run(cmd)
		fmt.Println(result)
	}
}

func TestShell_RunWithWait(t *testing.T) {
	s := getShell()
	tables := []string{
		"ping 127.0.0.1",
	}
	for _, cmd := range tables {
		result := s.Run(cmd)
		fmt.Println(result)
	}
}
