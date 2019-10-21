package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	gopath = strToolkit.Getrpath(os.Getenv("GOPATH"))
	if gopath == "" {
		gopath = strToolkit.Getrpath(fileToolkit.GetHomeDir())
	}

	modDir := gopath + "pkg/mod/"

	websiteList, e := listDirs(modDir)
	if e != nil {
		fmt.Println("list modDir error :", e)
		return
	}

	for _, website := range websiteList {
		if website == "cache" {
			continue
		}

		userList, e := listDirs(modDir + website)
		if e != nil {
			fmt.Println("list website error :", e)
			return
		}

		for _, user := range userList {
			if strings.HasPrefix(user, "!") {
				continue
			}

			repos, e := listDirs(modDir + website + sep + user)
			if e != nil {
				fmt.Println("list repo error :", e)
				return
			}

			for _, repoWithVersion := range repos {
				if strings.HasPrefix(repoWithVersion, ".") {
					continue
				}

				repo := getRepoName(repoWithVersion)
				relativeRepo, e := getRelativePath(website, user, repo)
				if e != nil {
					fmt.Println("get relative path error :", e)
					return
				}

				if relativeRepo == "" {
					continue
				}

				oldRepo := modDir + website + sep + user + sep + repoWithVersion
				e = copyDir(oldRepo, relativeRepo)
				if e != nil {
					fmt.Println("run cp -r error :", e)
					return
				}

				fmt.Println(website + sep + user + sep + repo)
			}
		}
	}
}

func copyDir(src, dst string) error {
	if runtime.GOOS == "windows" {
		return ioToolkit.RunAttachedCmd("robocopy", "/E", src, dst)
	}
	return ioToolkit.RunAttachedCmd("cp", "-r", src, dst)
}

func listDirs(root string) ([]string, error) {
	dirs := []string{}
	infos, e := ioutil.ReadDir(root)
	if e != nil {
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
	parentDir := gopath + "src" + sep + website + sep + user

	e := os.MkdirAll(parentDir, 0755)
	if e != nil {
		return "", errors.New("mkdirs failed:" + e.Error())
	}

	relativePath := parentDir + sep + repo
	_, e = os.Stat(relativePath)
	if os.IsNotExist(e) {
		return relativePath, nil
	}
	if e != nil {
		fmt.Println("stat file failed:", e)
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
