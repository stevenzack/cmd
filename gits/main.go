package main

import (
	"fmt"

	"github.com/stevenzack/cmd"
)

func main() {
	e := cmd.RunAttach("git", "status")
	if e != nil {
		fmt.Println("run error :", e)
		return
	}
}
