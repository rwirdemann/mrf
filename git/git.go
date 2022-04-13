package git

import "io/ioutil"

// FindRepositories finds all git repositories contained in the given directory.
func FindRepositories(dir string) map[int]string {
	m := make(map[int]string)
	files, _ := ioutil.ReadDir(dir)
	i := 1
	for _, file := range files {
		if file.IsDir() {
			m[i] = file.Name()
		}
		i++
	}
	return m
}
