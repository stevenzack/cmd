package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/stevenzack/cmd/tools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no input URL")
		return
	}

	url := os.Args[1]
	url = handleUrl(url)
	e := tools.RunAttach("git", "clone", "--depth=1", url)
	if e != nil {
		fmt.Println(e)
		return
	}
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
