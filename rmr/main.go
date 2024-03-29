package main

import (
	"fmt"
	"os"

	"github.com/stevenzack/cmd/internal/tools"
)

func main() {
	args := os.Args
	args[0] = "-Rf"
	e := tools.RunAttach("rm", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
