package main

import (
	"fmt"
	"os"

	"github.com/stevenzack/cmd/tools"
)

func main() {
	args := []string{"run", "-tags", "gtk_3_18", "main.go"}
	args = append(args, os.Args[1:]...)
	e := tools.RunAttach("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
