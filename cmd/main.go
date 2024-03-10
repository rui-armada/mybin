package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rui-armada/mybin/internal/git"
)

var (
	gitProvider git.GitProvider
	ctx         context.Context
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: mybin.exe <command> ...")
		return // fail
	}

	var err error
	command := args[1]
	ctx = context.Background()
	gitProvider, err = git.NewGithubProvider(ctx)
	if err != nil {
		fmt.Println("Failed to create GitHub provider: ", err)
		return // fail
	}

	switch command {
	case "create":
		handleCreate(args)
	case "delete":
		handleDelete(args)
	case "list":
		handleList()
	default:
		fmt.Println("Unknown command. Available commands: create, delete, list")
	}
}

func handleCreate(args []string) {
	if len(args) < 4 || args[2] != "--name" {
		fmt.Println("Usage: mybin.exe create --name <repo-name>")
		return
	}

	repoName := args[3]

	err := gitProvider.CreateRepo(ctx, repoName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleDelete(args []string) {
	if len(args) < 6 || args[2] != "--name" || args[4] != "--group" {
		fmt.Println("Usage: mybin.exe delete --name <repo-name> --group <group-name>")
		return
	}

	repoName := args[3]
	groupName := args[5]

	err := gitProvider.DeleteRepo(ctx, repoName, groupName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleList() {
	err := gitProvider.ListRepos(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}
