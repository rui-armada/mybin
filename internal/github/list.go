package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func ListRepos() error {
	// Retrieve the GitHub personal access token from an environment variable
	token := os.Getenv("github_token")
	if token == "" {
		return fmt.Errorf("GitHub token not found")
	}

	// Create an authenticated GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// List repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return fmt.Errorf("Failed to list repositories: %s", err)
	}

	fmt.Println("Your repositories:")
	for _, repo := range repos {
		fmt.Println(*repo.FullName)
	}

	return nil
}
