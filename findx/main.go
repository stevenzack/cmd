package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/tools/fileToolkit"
)

var (
	f       = flag.String("f", "", "Specific the file to find")
	c       = flag.Bool("c", false, "Find all lines with Chinese")
	exclude = flag.String("e", "", "Exclude file")
	v       = flag.Bool("v", false, "Verbose on every file")
	//
	wd string
)

func main() {
	if *f != "" {
		handleFile(*f)
		return
	}
	var e error
	wd, e = os.Getwd()
	if e != nil {
		log.Fatal(e)
	}

	e = filepath.Walk(wd, func(path string, info os.FileInfo, e error) error {

		ext := filepath.Ext(path)
		switch ext {
		case ".go", ".swift", ".java", ".txt", ".js", ".html", ".css", ".xml":
		default:
			return nil
		}
		if *exclude != "" {
			s1, e := filepath.Abs(path)
			if e != nil {
				log.Println(e)
				return e
			}
			s2, e := filepath.Abs(*exclude)
			if e != nil {
				log.Println(e)
				return e
			}
			if s1 == s2 {
				return nil
			}
		}

		handleFile(path)
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}

func handleFile(path string) {
	content, e := fileToolkit.ReadFileAll(path)
	if e != nil {
		log.Fatal(e)
	}
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strToolkit.HasChinese(line) {
			fmt.Println(path[len(wd)+1:]+":"+strconv.Itoa(i+1), ":", strToolkit.TrimBoth(strToolkit.TrimBoth(line, " "), "\t"))
			if !*v {
				return
			}
		}
	}
}
