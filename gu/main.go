package main

import (
	"github.com/StevenZack/tools/cmdToolkit"
	"github.com/StevenZack/tools/ioToolkit"

	// "bufio"
	"flag"
	"fmt"
)

var branch = flag.String("b", "master", "branch")
var tag = flag.String("t", "", "Add tag")

func main() {
	flag.Parse()
	_, e := cmdToolkit.Run("git", "add", "--all")
	if e != nil {
		fmt.Println("git", "add:", e)
		return
	}
	m := "Just a Backup"
	if len(flag.Args()) > 0 {
		m = flag.Arg(0)
	}
	_, e = cmdToolkit.Run("git", "commit", "-m", m)
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

	_, e = cmdToolkit.Run("git", "push")
	if e != nil {
		fmt.Println(e)
		return
	}
	_, e = cmdToolkit.Run("git", "push", "--tags")
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
