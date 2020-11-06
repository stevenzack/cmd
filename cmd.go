package cmd

import (
	"os"
	"os/exec"
)

func RunAttach(p string, args ...string) error {
	c := exec.Command(p, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
