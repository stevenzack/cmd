package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("go", "run", "-tags", "gtk_3_18", "main.go")
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	e := c.Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
