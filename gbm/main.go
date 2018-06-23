package main

import (
	"os/exec"
	"os"
	"fmt"
)

func main() {
	c:=exec.Command("go","build","-tags","gtk_3_18","main.go")
	c.Stderr=os.Stderr
	c.Stdout=os.Stdout
	c.Stdin=os.Stdin
	e:=c.Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}