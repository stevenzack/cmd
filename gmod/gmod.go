package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/StevenZack/cmd"
	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

var (
	gopath = ""
)

func main() {
	log.SetFlags(log.Lshortfile)

	gopath = strToolkit.Getrpath(os.Getenv("GOPATH"))
	if gopath == "" {
		gopath = filepath.Join(fileToolkit.GetHomeDir(), "go")
	}

	modDir := filepath.Join(gopath, "pkg/mod/")

	websiteList, e := listDirs(modDir)
	if e != nil {
		log.Println(e)
		return
	}

	for _, website := range websiteList {
		if website == "cache" {
			continue
		}

		userList, e := listDirs(filepath.Join(modDir, website))
		if e != nil {
			log.Println(e)
			return
		}

		for _, user := range userList {

			repos, e := listDirs(filepath.Join(modDir, website, user))
			if e != nil {
				log.Println(e)
				return
			}

			for _, repoWithVersion := range repos {
				if strings.HasPrefix(repoWithVersion, ".") {
					continue
				}

				repo := getRepoName(repoWithVersion)
				relativeRepo, e := getRelativePath(website, user, repo)
				if e != nil {
					log.Println(e)
					return
				}

				if relativeRepo == "" {
					continue
				}

				// oldRepo := modDir + website + sep + user + sep + repoWithVersion
				oldRepo := filepath.Join(modDir, website, user, repoWithVersion)
				e = copyDir(oldRepo, relativeRepo)
				if e != nil {
					log.Println(e)
					return
				}

				fmt.Println(filepath.Join(website, user, repo))
			}
		}
	}
}

func translateUpperCase(s string) string {
	s = strToolkit.SubBefore(s, "@v", s)
	rp := ""
	upperCase := false
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "!" {
			upperCase = true
			continue
		}
		v := s[i : i+1]
		if upperCase {
			v = strings.ToUpper(v)
			upperCase = false
		}
		rp += v
	}
	return rp
}

func copyDir(src, dst string) error {
	if runtime.GOOS == "windows" {
		exec.Command("robocopy", "/E", src, dst).Run()
		return nil
	}
	return cmd.RunAttach("cp", "-r", src, dst)
}

func listDirs(root string) ([]string, error) {
	dirs := []string{}
	infos, e := ioutil.ReadDir(root)
	if e != nil {
		log.Println(e)
		return nil, e
	}

	for _, info := range infos {
		if info.IsDir() {
			dirs = append(dirs, info.Name())
		}
	}
	return dirs, nil
}

func getRelativePath(website, user, repo string) (string, error) {
	parentDir := filepath.Join(gopath, "src", translateUpperCase(website), translateUpperCase(user))

	e := os.MkdirAll(parentDir, 0755)
	if e != nil {
		return "", errors.New("mkdirs failed:" + e.Error())
	}

	relativePath := filepath.Join(parentDir, translateUpperCase(repo))
	_, e = os.Stat(relativePath)
	if os.IsNotExist(e) {
		return relativePath, nil
	}
	if e != nil {
		log.Println(e)
		return "", e
	}

	return "", nil
}

func getRepoName(repo string) string {
	i := strings.Index(repo, "@")
	if i == -1 {
		return repo
	}
	return repo[:i]
}
