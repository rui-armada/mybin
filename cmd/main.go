package main

import (
	"fmt"
	"os"

	"github.com/yourusername/github-cli/internal/github"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: mybin.exe <command> ...")
		return
	}

	command := args[1]

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

	err := github.CreateRepo(repoName)
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

	err := github.DeleteRepo(repoName, groupName)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleList() {
	err := github.ListRepos()
	if err != nil {
		fmt.Println(err)
		return
	}
}
