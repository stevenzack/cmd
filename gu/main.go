package main

import (
	"github.com/StevenZack/cmd"
	"github.com/StevenZack/tools/ioToolkit"

	// "bufio"
	"flag"
	"fmt"
)

var branch = flag.String("b", "master", "branch")
var tag = flag.String("t", "", "Add tag")

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
		fmt.Println(e)
		return
	}

	//tag
	if *tag != "" {
		e = ioToolkit.RunAttachedCmd("git", "tag", "-a", *tag, "-m", m)
		if e != nil {
			fmt.Println("add tag error :", e)
			return
		}
	}

	e = cmd.NewCmd("git", "push").Run()
	if e != nil {
		fmt.Println(e)
		return
	}
	e = cmd.NewCmd("git", "push", "--tags").Run()
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
