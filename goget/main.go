package main

import (
	"fmt"
	"github.com/StevenZack/tools/strToolkit"
	"os"
	"os/exec"
)

func main() {
	url := os.Args[1]
	url = handleUrl(url)
	cmd := exec.Command("go", "get", url)
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
	if strToolkit.StartsWith(s, preffix) {
		s = s[len(preffix):]
	}
	if strToolkit.EndsWith(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
