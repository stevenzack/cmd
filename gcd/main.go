package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/stevenzack/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no input URL")
		return
	}

	url := os.Args[1]
	url = handleUrl(url)
	e := cmd.RunAttach("git", "clone", "--depth=1", url)
	if e != nil {
		fmt.Println(e)
		return
	}
}
func countStr(str, s string) int {
	c := 0
	for i := 0; i < len(str); i++ {
		if str[i:i+1] == s {
			c++
		}
	}
	return c
}
func handleUrl(s string) string {
	url := s
	prefix := "https://"
	prefix2 := "http://"
	if strings.HasPrefix(url, prefix) || strings.HasPrefix(s, prefix2) {
		return url
	}
	return prefix + url
}
