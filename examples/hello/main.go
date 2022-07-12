package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_shell"
)

func main() {
	s := zdpgo_shell.New()

	tables := []string{
		"dir",
		"cd",
	}
	for _, cmd := range tables {
		result := s.Run(cmd)
		if result.Error != "" {
			fmt.Println(result.Error)
		}
		fmt.Println(result)
	}
}
