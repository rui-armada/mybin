package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/rui-armada/mybin/internal/git"
	"github.com/spf13/cobra"
)

var (
	gitProvider git.GitProvider
	ctx         context.Context
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mybin",
	Short: "MyBin is a CLI tool to manage your repositories.",
	Long: `MyBin is a CLI tool to manage your repositories on your git provider.
It allows you to create, delete, ... repositories with extra functionalities.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var err error
	ctx = context.Background()
	gitProvider, err = git.NewGithubProvider(ctx)
	if err != nil {
		fmt.Println("Failed to create GitHub provider: ", err)
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mybin.yaml)")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
