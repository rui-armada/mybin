package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func CreateRepo(repoName string) error {
	// Retrieve the GitHub personal access token from an environment variable
	token := "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"
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

	// Create the repository
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(false),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return fmt.Errorf("Failed to create repository: %s", err)
	}

	fmt.Printf("Repository '%s' created successfully!\n", repoName)
	return nil
}
