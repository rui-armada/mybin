package git

import "context"

type GitProvider interface {
	// CreateRepo creates a new repository
	CreateRepo(ctx context.Context, repoName string) error

	// DeleteRepo deletes a repository
	DeleteRepo(ctx context.Context, repoName string, groupName string) error

	// ListRepos lists all repositories
	ListRepos(ctx context.Context) error
}
