package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func DeleteRepo(repoName string, groupName string) error {
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

	// Delete the repository
	_, err := client.Repositories.Delete(ctx, groupName, repoName)
	if err != nil {
		return fmt.Errorf("Failed to delete repository: %s", err)
	}

	fmt.Printf("Repository '%s' deleted successfully!\n", repoName)
	return nil
}
