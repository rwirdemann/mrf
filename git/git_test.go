package git

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGetRepositories(t *testing.T) {
	baseDir := removeAndCreateBaseDir()

	mkDir(baseDir, "repo")
	mkDir(baseDir, "repo/.git")
	mkDir(baseDir, "norepo")

	repositories := FindRepositories(baseDir)
	assert.True(t, contains(repositories, "repo"))
	assert.False(t, contains(repositories, "norepo"))

	err := os.RemoveAll(baseDir)
	if err != nil {
		log.Fatal(err)
	}
}

func contains(repositories map[int]string, repo string) bool {
	for _, v := range repositories {
		if v == repo {
			return true
		}
	}
	return false
}

func removeAndCreateBaseDir() string {
	tempDir := os.TempDir()
	if tempDir == "" {
		tempDir = "."
	}

	baseDir := fmt.Sprintf("%s/testdata", tempDir)
	err := os.RemoveAll(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(baseDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	return baseDir
}

func mkDir(base string, dir string) {
	repo := fmt.Sprintf("%s/%s", base, dir)
	err := os.Mkdir(repo, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
