package zdpgo_shell

import (
	"runtime"
)

/*
@Time : 2022/4/29 21:45
@Author : 张大鹏
@File : shell
@Software: Goland2021.3.1
@Description: 核心对象和方法
*/

type Shell struct {
	Config *Config // 配置对象
}

func New() *Shell {
	return NewWithConfig(Config{})
}

func NewWithConfig(config Config) *Shell {
	s := Shell{}

	// 配置
	if config.ShellPath == "" {
		switch runtime.GOOS {
		case "linux":
			config.ShellPath = "/bin/bash"
		case "windows":
			config.ShellPath = "C:\\Windows\\System32\\cmd.exe"
		default:
			config.ShellPath = "/bin/sh"
		}
	}
	s.Config = &config
	return &s
}

//Run 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func (s *Shell) Run(command string) string {
	result := RunCommand(command)
	return result
}

// RunWithError 运行命令并返回结果和错误
func (s *Shell) RunWithError(command string) (string, error) {
	result, err := RunCommandWithErr(command)
	return result, err
}

// RunWithWait 执行命令并等待结果返回，适合ping之类的需要等待一段时间的命令
func (s *Shell) RunWithWait(command string) (string, error) {
	result, err := RunWithWait(command)
	return result, err
}
