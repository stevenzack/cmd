package main

import (
	"fmt"
	"os"

	"github.com/stevenzack/cmd"
)

func main() {
	args := os.Args
	args[0] = "-Rf"
	e := cmd.RunAttach("rm", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
