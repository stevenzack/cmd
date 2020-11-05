package main

import (
	"fmt"
	"os"

	"github.com/StevenZack/tools/cmdToolkit"
)

func main() {
	args := append([]string{"install", "-tags", "gtk_3_18"}, os.Args[1:]...)
	_, e := cmdToolkit.Run("go", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
