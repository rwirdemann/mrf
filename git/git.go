package git

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FindRepositories finds all git repositories contained in the given directory.
func FindRepositories(dir string) map[int]string {
	m := make(map[int]string)
	files, _ := ioutil.ReadDir(dir)
	i := 1
	for _, file := range files {
		if file.IsDir() && isGitDir(fmt.Sprintf("%s/%s", dir, file.Name())) {
			m[i] = file.Name()
		}
		i++
	}
	return m
}

func isGitDir(file string) bool {
	fileInfo, err := os.Stat(fmt.Sprintf("%s/.git", file))
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
