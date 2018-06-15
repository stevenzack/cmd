package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	url := os.Args[1]
	url = handleUrl(url)
	cmd := exec.Command("go", "install", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
func handleUrl(url string) string {
	s := url
	preffix, suffix := "https://", ".git"
	if len(s) > len(preffix) && s[:len(preffix)] == preffix {
		s = s[len(preffix):]
	}
	if len(s) > len(suffix) && s[len(s)-len(suffix):] == suffix {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
