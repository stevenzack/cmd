package main

import (
	"fmt"
	"github.com/StevenZack/cmd"
	"os"
)

func main() {
	e := cmd.NewCmd("gw").Run()
	if e != nil {
		fmt.Println(`gw error:`, e)
		return
	}
	if len(os.Args) > 1 {
		e = cmd.NewCmd("go", "run", os.Args[1]).Run()
		if e != nil {
			fmt.Println(`run error:`, e)
			return
		}
	}
}
