package main

import (
	"fmt"
	"mime"
	"os"

	"github.com/stevenzack/tools/strToolkit"
)

func main() {
	fmt.Println(os.Args[1])
	fmt.Println(mime.TypeByExtension("." + strToolkit.TrimStart(os.Args[1], ".")))
}
