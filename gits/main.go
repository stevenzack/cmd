package main

import (
	"fmt"

	"github.com/StevenZack/tools/ioToolkit"
)

func main() {
	e := ioToolkit.RunAttachedCmd("git", "status")
	if e != nil {
		fmt.Println("run error :", e)
		return
	}
}
