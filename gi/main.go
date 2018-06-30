package main

import (
	"fmt"
	"github.com/StevenZack/cmd"
	"os"
)

func main() {
	args := append([]string{"install", "-tags", "gtk_3_18"}, os.Args[1:]...)
	e := cmd.NewCmd("go", args...).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
