package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	e := exec.Command("git", "add", "--all").Run()
	if e != nil {
		fmt.Println("git", "add:", e)
		return
	}
	m := "Just a Backup"
	if len(os.Args) > 1 {
		m = os.Args[1]
	}
	e = exec.Command("git", "commit", "-m", m).Run()
	if e != nil {
		if e.Error() != "exit status 128" {
			fmt.Println("git", "commit:", e)
			return
		} else {
			e = exec.Command("git", "config", "--global", "user.email", "stevenzack@qq.com").Run()
			if e != nil {
				fmt.Println("git", "config", "--global", "user.email:", e)
				return
			}
			e = exec.Command("git", "config", "--global", "user.name", "StevenZack").Run()
			if e != nil {
				fmt.Println("git", "config", "--global", "user.name:", e)
				return
			}
			e = exec.Command("git", "config", "--global", "credential.helper", "store").Run()
			if e != nil {
				fmt.Println("credential.helper:", e)
				return
			}
		}
	}
	cmd := exec.Command("git", "push", "origin", "master")
	stdout, e := cmd.StderrPipe()
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
	}
	cmd.Wait()
}
