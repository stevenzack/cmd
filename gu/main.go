package main

import (
	"flag"
	"fmt"

	"github.com/StevenZack/cmd"
)

var branch = flag.String("b", "master", "branch")
var tag = flag.String("t", "", "Add tag")

func main() {
	flag.Parse()
	e := cmd.RunAttach("git", "add", "--all")
	if e != nil {
		fmt.Println("git", "add:", e)
		return
	}
	m := "Just a Backup"
	if len(flag.Args()) > 0 {
		m = flag.Arg(0)
	}
	e = cmd.RunAttach("git", "commit", "-m", m)
	if e != nil {
		fmt.Println(e)
		return
	}

	//tag
	if *tag != "" {
		e = cmd.RunAttach("git", "tag", "-a", *tag, "-m", m)
		if e != nil {
			fmt.Println("add tag error :", e)
			return
		}
	}

	e = cmd.RunAttach("git", "push")
	if e != nil {
		fmt.Println(e)
		return
	}
	e = cmd.RunAttach("git", "push", "--tags")
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
