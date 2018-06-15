package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	args[0] = "-Rf"
	e := exec.Command("rm", args...).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
