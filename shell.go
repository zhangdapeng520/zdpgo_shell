package zdpgo_shell

import (
	"github.com/zhangdapeng520/zdpgo_daemon"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

	s.Config = &config
	return &s
}

//Run 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func (s *Shell) Run(command string) Result {
	result := Run(command)
	return result
}

// RunRealTime 执行命令并等待结果返回，适合ping之类的需要等待一段时间的命令
func (s *Shell) RunRealTime(command string) Result {
	result := RunRealTime(command)
	return result
}

// RunBackground 开启新的进程执行
func (s *Shell) RunBackground(executeFile string, args ...string) Result {
	// 启动一个子进程后主程序退出
	zdpgo_daemon.Background("", false)

	// 写入进程ID
	ioutil.WriteFile(".env", []byte(strconv.Itoa(os.Getpid())), os.ModePerm)

	// 以下代码只有子程序会执行
	log.Println(os.Getpid(), "start...")
	result := s.Run(executeFile)
	log.Println(os.Getpid(), "end")

	// 返回
	return result
}
