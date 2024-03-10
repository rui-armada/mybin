package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	deleteRepoName  string
	deleteGroupName string
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteRepoName == "" {
			fmt.Println("'--repo-name' flag is mandatory and cannot be empty.")
			return
		}

		if deleteGroupName == "" {
			fmt.Println("'--group-name' flag is mandatory and cannot be empty.")
			return
		}

		err := gitProvider.DeleteRepo(ctx, deleteRepoName, deleteGroupName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&deleteRepoName, "repo-name", "r", "", "Name of the repository to delete")
	deleteCmd.Flags().StringVarP(&deleteGroupName, "group-name", "g", "", "Name of the group")
}
