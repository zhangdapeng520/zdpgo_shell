package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_shell"
)

func main() {
	s := zdpgo_shell.New()
	result := s.RunBackground(".\\shell.exe")
	fmt.Println(result)
}
