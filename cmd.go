package cmd

import (
	"os"
	"os/exec"
)

func NewCmd(name string, args ...string) *exec.Cmd {
	c := exec.Command(name, args...)
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	return c
}
