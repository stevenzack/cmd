package main

import (
	"fmt"
	"os"

	"github.com/stevenzack/cmd/tools"
)

func main() {
	args := append([]string{"install", "-tags", "gtk_3_18"}, os.Args[1:]...)
	e := tools.RunAttach("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
