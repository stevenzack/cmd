package main

import (
	"fmt"
	"os"

	"github.com/stevenzack/cmd/internal/tools"
)

func main() {
	args := append([]string{"install"}, os.Args[1:]...)
	e := tools.RunAttach("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
