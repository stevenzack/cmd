package tools

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunAttach(cmd string, args ...string) error {
	fmt.Println(cmd, strings.Join(args, " "))
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
