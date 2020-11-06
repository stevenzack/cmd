package main

import (
	"fmt"

	"github.com/StevenZack/cmd"
)

func main() {
	e := cmd.RunAttach("git", "status")
	if e != nil {
		fmt.Println("run error :", e)
		return
	}
}
