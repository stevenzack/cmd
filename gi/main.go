package main

import (
	"fmt"
	"os"

	"github.com/StevenZack/cmd"
)

func main() {
	args := append([]string{"install", "-tags", "gtk_3_18"}, os.Args[1:]...)
	e := cmd.RunAttach("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
