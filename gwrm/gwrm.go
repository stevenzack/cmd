package main

import (
	"fmt"

	"github.com/StevenZack/cmd"
)

func main() {
	e := cmd.NewCmd("gwr", "main.go").Run()
	if e != nil {
		fmt.Println(`run error:`, e)
		return
	}
}
