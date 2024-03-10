package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createRepoName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if createRepoName == "" {
			fmt.Println("'--repo-name' flag is mandatory and cannot be empty.")
			return
		}

		err := gitProvider.CreateRepo(ctx, createRepoName)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&createRepoName, "repo-name", "r", "", "Name of the repository to create")
}
