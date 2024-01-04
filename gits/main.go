package main

import (
	"fmt"

	"github.com/stevenzack/cmd/tools"
)

func main() {
	e := tools.RunAttach("git", "status")
	if e != nil {
		fmt.Println("run error :", e)
		return
	}
}
