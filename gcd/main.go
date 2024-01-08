package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/stevenzack/cmd/internal/tools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no input URL")
		return
	}

	url := os.Args[1]
	url = handleUrl(url)
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	if !containsDepth(args) {
		args = append(args, "--depth=1")
	}
	args = append([]string{"clone", url}, args...)
	e := tools.RunAttach("git", args...)
	if e != nil {
		fmt.Println(e)
		return
	}
}
func containsDepth(args []string) bool {
	for _, v := range args {
		if strings.Contains(v, "depth") {
			return true
		}
	}
	return false
}
func handleUrl(s string) string {
	url := s
	prefix := "https://"
	prefix2 := "http://"
	if strings.HasPrefix(url, prefix) || strings.HasPrefix(s, prefix2) || strings.HasPrefix(s, "git@") {
		return url
	}
	if strings.Count(s, "/") == 1 {
		return "git@github.com:" + s + ".git"
	}
	return prefix + url
}
