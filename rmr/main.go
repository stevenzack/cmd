package main

import (
	"fmt"
	"os"

	"github.com/StevenZack/tools/cmdToolkit"
)

func main() {
	args := os.Args
	args[0] = "-Rf"
	_, e := cmdToolkit.Run("rm", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
