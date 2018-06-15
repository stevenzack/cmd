package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no enough args")
		return
	}
	args := os.Args
	args[0] = "-w"
	c := exec.Command("goimports", args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	e := c.Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
