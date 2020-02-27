package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/ioToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

var (
	sep    = string(os.PathSeparator)
	gopath = ""
)

func main() {
	log.SetFlags(log.Lshortfile)

	gopath = strToolkit.Getrpath(os.Getenv("GOPATH"))
	if gopath == "" {
		gopath = strToolkit.Getrpath(fileToolkit.GetHomeDir())
	}

	modDir := gopath + "pkg/mod/"

	websiteList, e := listDirs(modDir)
	if e != nil {
		log.Println(e)
		return
	}

	for _, website := range websiteList {
		if website == "cache" {
			continue
		}

		userList, e := listDirs(modDir + website)
		if e != nil {
			log.Println(e)
			return
		}

		for _, user := range userList {

			repos, e := listDirs(modDir + website + sep + user)
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

				oldRepo := modDir + website + sep + user + sep + repoWithVersion
				e = copyDir(oldRepo, relativeRepo)
				if e != nil {
					log.Println(e)
					return
				}

				fmt.Println(website + sep + user + sep + repo)
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
			v = strToolkit.ToUpper(v)
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
	return ioToolkit.RunAttachedCmd("cp", "-r", src, dst)
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
	parentDir := gopath + "src" + sep + translateUpperCase(website) + sep + translateUpperCase(user)

	e := os.MkdirAll(parentDir, 0755)
	if e != nil {
		return "", errors.New("mkdirs failed:" + e.Error())
	}

	relativePath := parentDir + sep + translateUpperCase(repo)
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
