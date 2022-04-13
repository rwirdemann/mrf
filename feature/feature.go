package feature

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func Start(featureName string, repository string) {
	os.Chdir(fmt.Sprintf("/Users/ralf/Documents/work/betzemeier/bm-app-suite/%s", repository))
	log.Printf("Starting feature %s", featureName)
	mvn("gitflow:feature-start", featureName)
}

func Finish(featureName string, repository string) {
	os.Chdir(fmt.Sprintf("/Users/ralf/Documents/work/betzemeier/bm-app-suite/%s", repository))
	log.Printf("Finishing feature %s", featureName)
	mvn("gitflow:feature-finish", featureName)
}

func mvn(goal string, featureName string) {
	cmd := exec.Command("mvn", "-U", "-B", "-DfetchRemote=false", goal, fmt.Sprintf("-DfeatureName=%s", featureName))

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %goal\n", err)
	}
	outStr, _ := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%goal\n", outStr)
}
