package main

import (
	"fmt"
	"os"

	"github.com/stevenzack/cmd"
)

func main() {
	args := []string{"run", "-tags", "gtk_3_18", "main.go"}
	args = append(args, os.Args[1:]...)
	e := cmd.RunAttach("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
