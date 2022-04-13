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

// featureStartCmd represents the feature-start command
var featureStartCmd = &cobra.Command{
	Use:   "feature-start",
	Short: "Starts a feature branch and optionally updates version(s).",
	Long:  `Starts a feature branch and optionally updates version(s).`,
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
		fmt.Printf("Creating a new branch 'feature/%s' on repo(s): %s\n", branch, affectedRepositories)
		r := strings.Split(affectedRepositories, ",")
		for _, v := range r {
			i, _ := strconv.Atoi(v)
			repo := repositories[i]
			fmt.Printf("Processing '%s'\n", repo)
			feature.Start(branch, repo)
		}
		os.Chdir(currentDir)
	},
}

func init() {
	rootCmd.AddCommand(featureStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// featureStartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// featureStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
