package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/myapp/feature"
	"github.com/spf13/myapp/git"
	"os"
	"strconv"
	"strings"
)

// featureFinishCmd represents the feature-finish command
var featureFinishCmd = &cobra.Command{
	Use:   "feature-finish",
	Short: "Merges a feature branch.",
	Long:  `Merges a feature branch back to develop.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("What is a name of feature branch? feature/:")
		var branch string
		fmt.Scanln(&branch)

		currentDir, _ := os.Getwd()

		workingDir := os.Getenv("WORKING_DIR")
		os.Chdir(workingDir)
		repositories := git.FindRepositories(workingDir)
		for k, v := range repositories {
			fmt.Printf("%d: %s\n", k, v)
		}
		fmt.Print("Select affected affected repositories (e.g. 1,2):")
		var affectedRepositories string
		fmt.Scanln(&affectedRepositories)
		fmt.Printf("Finishing branch 'feature/%s' on repo(s): %s\n", branch, affectedRepositories)
		r := strings.Split(affectedRepositories, ",")
		for _, v := range r {
			i, _ := strconv.Atoi(v)
			repo := repositories[i]
			fmt.Printf("Processing '%s'\n", repo)
			feature.Finish(branch, repo)
		}
		os.Chdir(currentDir)
	},
}

func init() {
	rootCmd.AddCommand(featureFinishCmd)
}
