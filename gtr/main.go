package main

import (
	"log"
	"os"

	"github.com/stevenzack/cmd/tools"
)

func main() {
	args := []string{}
	if len(os.Args) > 1 {
		args = append(args, "-run="+os.Args[1])
	}
	e := tools.RunAttach("go", append([]string{"test", "-v"}, args...)...)
	if e != nil {
		log.Fatal(e)
	}
}
