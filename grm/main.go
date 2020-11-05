package main

import (
	"fmt"
	"os"

	"github.com/StevenZack/tools/cmdToolkit"
)

func main() {
	args := []string{"run", "-tags", "gtk_3_18", "main.go"}
	args = append(args, os.Args[1:]...)
	_, e := cmdToolkit.Run("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
