package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stevenzack/cmd/internal/tools"
)

var p = flag.Bool("p", false, "git pull")

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	flag.Parse()
	arg := "status"
	if *p {
		arg = "pull"
	}
	pwd, e := os.Getwd()
	if e != nil {
		log.Panic(e)
		return
	}
	const divider = "------------"
	if _, e := os.Stat(".git"); os.IsNotExist(e) {
		list, e := os.ReadDir(".")
		if e != nil {
			log.Println(e)
			return
		}
		for _, item := range list {
			if _, e := os.Stat(item.Name() + "/.git"); os.IsNotExist(e) {
				continue
			}
			e = os.Chdir(item.Name())
			if e != nil {
				log.Println(e, item.Name())
				continue
			}
			println(divider, item.Name(), divider)
			e = tools.RunAttach("git", arg)
			if e != nil {
				log.Println(e)
				return
			}
			e = os.Chdir(pwd)
			if e != nil {
				log.Println(e, pwd)
				continue
			}
		}
		return
	}

	if _, e := os.Stat("go.mod"); e == nil {
		b, e := os.ReadFile("go.mod")
		if e != nil {
			log.Println(e)
			return
		}

		for _, line := range strings.Split(string(b), "\n") {
			if strings.Contains(line, "replace") {
				line = tools.SubAfterLast(line, "=>", "")
				line = strings.TrimSpace(line)
				e = os.Chdir(line)
				if e != nil {
					log.Println(e, line)
					continue
				}

				println(divider, line, divider)
				e = tools.RunAttach("git", arg)
				if e != nil {
					log.Println(e)
					return
				}
			}
		}
		os.Chdir(pwd)
	}

	println(divider, ".", divider)
	e = tools.RunAttach("git", arg)
	if e != nil {
		fmt.Println("run error :", e)
		return
	}
}
