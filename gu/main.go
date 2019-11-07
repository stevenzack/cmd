package main

import (
	"github.com/StevenZack/cmd"
	// "bufio"
	"flag"
	"fmt"
)

var branch = flag.String("b", "master", "branch")

func main() {
	flag.Parse()
	e := cmd.NewCmd("git", "add", "--all").Run()
	if e != nil {
		fmt.Println("git", "add:", e)
		return
	}
	m := "Just a Backup"
	if len(flag.Args()) > 0 {
		m = flag.Arg(0)
	}
	e = cmd.NewCmd("git", "commit", "-m", m).Run()
	if e != nil {
		if e.Error() == "exit status 1" {
			fmt.Println("nothing to commit")
			return
		} else if e.Error() != "exit status 128" {
			fmt.Println("git", "commit:", e)
			return
		} else {
			e = cmd.NewCmd("git", "config", "--global", "user.email", "stevenzack@qq.com").Run()
			if e != nil {
				fmt.Println("git", "config", "--global", "user.email:", e)
				return
			}
			e = cmd.NewCmd("git", "config", "--global", "user.name", "StevenZack").Run()
			if e != nil {
				fmt.Println("git", "config", "--global", "user.name:", e)
				return
			}
			e = cmd.NewCmd("git", "config", "--global", "credential.helper", "store").Run()
			if e != nil {
				fmt.Println("credential.helper:", e)
				return
			}
		}
	}
	e = cmd.NewCmd("git", "push").Run()
	if e != nil {
		fmt.Println(e)
		return
	}
	// stdout, e := cmd.StderrPipe()
	// cmd.Start()
	// scanner := bufio.NewScanner(stdout)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	txt := scanner.Text()
	// 	fmt.Print(txt)
	// }
	// cmd.Wait()
}
