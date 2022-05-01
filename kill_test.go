package zdpgo_shell

import "testing"

func TestShell_Kill(t *testing.T) {
	s := getShell()
	s.Kill(3940)
}
