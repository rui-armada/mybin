package git

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

type githubProvider struct {
	*github.Client
}

func NewGithubProvider(ctx context.Context) (GitProvider, error) {
	// Retrieve the GitHub personal access token from an environment variable
	token := os.Getenv("github_token")
	if token == "" {
		return nil, fmt.Errorf("GitHub token not found")
	}

	// Create an authenticated GitHub client
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
	client := github.NewClient(tc)

	return &githubProvider{client}, nil
}

func (g *githubProvider) CreateRepo(ctx context.Context, repoName string) error {
	// Create the repository
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(false),
	}
	_, _, err := g.Client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return fmt.Errorf("Failed to create repository: %s", err)
	}

	fmt.Printf("Repository '%s' created successfully!\n", repoName)
	return nil
}

func (g *githubProvider) DeleteRepo(ctx context.Context, repoName string, groupName string) error {
	// Delete the repository
	_, err := g.Client.Repositories.Delete(ctx, groupName, repoName)
	if err != nil {
		return fmt.Errorf("Failed to delete repository: %s", err)
	}

	fmt.Printf("Repository '%s' deleted successfully!\n", repoName)
	return nil
}

func (g *githubProvider) ListRepos(ctx context.Context) error {
	// List repositories for the authenticated user
	repos, _, err := g.Client.Repositories.List(ctx, "", nil)
	if err != nil {
		return fmt.Errorf("Failed to list repositories: %s", err)
	}

	fmt.Println("Your repositories:")
	for _, repo := range repos {
		fmt.Println(*repo.FullName)
	}

	return nil
}
