package main

import (
	"fmt"
	"github.com/StevenZack/cmd"
	"os"
)

func main() {
	args := os.Args
	args[0] = "-Rf"
	e := cmd.NewCmd("rm", args...).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
