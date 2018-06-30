package main

import (
	"fmt"
	"github.com/StevenZack/cmd"
	"runtime"
)

func main() {
	e := cmd.NewCmd("go", "build", "-tags", "gtk_3_18", "main.go").Run()
	if e != nil {
		fmt.Println(e)
		return
	}
	name := "main.exe"
	if runtime.GOOS != "windows" {
		name = "./main"
	}
	e = cmd.NewCmd(name).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
