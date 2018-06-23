package main

import (
	"os/exec"
	"os"
	"fmt"
	"runtime"
)

func main() {
	e:=newCommand("go","build","-tags","gtk_3_18","main.go").Run()
	if e != nil {
		fmt.Println(e)
		return
	}
	name:="main.exe"
	if runtime.GOOS!="windows" {
		name="./main"
	}
	e=newCommand(name).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
func newCommand(name string,arg ...string)*exec.Cmd  {
	c:=exec.Command(name,arg...)
	c.Stderr=os.Stderr
	c.Stdin=os.Stdin
	c.Stdout=os.Stdout
	return c
}