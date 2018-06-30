package main

import (
	"fmt"
	"github.com/StevenZack/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no enough args")
		return
	}
	args := os.Args
	args[0] = "-w"
	e := cmd.NewCmd("goimports", args...).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
