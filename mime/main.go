package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	a := filepath.Ext(os.Args[1])
	fmt.Println(a)
	fmt.Println(mime.TypeByExtension("." + strings.TrimPrefix(a, ".")))
}
