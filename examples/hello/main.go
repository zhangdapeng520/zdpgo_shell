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
		result, err := s.RunWithError(cmd)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}
