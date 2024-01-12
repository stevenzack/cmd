package main

import (
	"fmt"
	"os"

	"github.com/stevenzack/cmd/internal/tools"
)

func main() {
	args := []string{"run", "main.go"}
	args = append(args, os.Args[1:]...)
	e := tools.RunAttach("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
