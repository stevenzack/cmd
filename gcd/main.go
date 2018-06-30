package main

import (
	"fmt"
	"github.com/StevenZack/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no input URL")
		return
	}

	url := os.Args[1]
	url = handleUrl(url)
	e := cmd.NewCmd("git", "clone", "--depth=1", url).Run()
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
	preffix := "https://"
	if len(url) > len(preffix) && url[:len(preffix)] != preffix || len(url) <= len(preffix) {
		url = preffix + url
	}
	return url
}
