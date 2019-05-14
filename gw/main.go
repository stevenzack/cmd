package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"github.com/StevenZack/cmd"
)

func main() {
	if len(os.Args) < 2 {
		ls,e:=ioutil.ReadDir(".")
		if e!=nil{
		fmt.Println(`read dir error:`,e)
		return
		}
		for _,l:=range ls{
			if strings.HasSuffix(l.Name(), ".go"){
				e:=cmd.NewCmd("goimports", "-w",l.Name()).Run()
				if e!=nil{
				fmt.Println(`run cmd error:`,e)
				return
				}
			}
		}
		return
	}
	args := os.Args
	args[0] = "-w"
	e := cmd.NewCmd("goimports", args...).Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
