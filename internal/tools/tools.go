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

func SubAfterLast(s, sep, def string) string {
	for i := len(s) - len(sep); i > -1; i-- {
		v := s[i : i+len(sep)]
		if v == sep {
			return s[i+len(sep):]
		}
	}
	return def
}
