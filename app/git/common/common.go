package common

import (
	"github.com/go-git/go-git/v5"
	"log"
)

type Git struct {
	Repo     *git.Repository
	Worktree *git.Worktree
	Remotes  []*git.Remote
}

func GetRepo(path string) *git.Repository {
	// Open the repository
	repository, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	return repository
}

func GetWorktree(repository *git.Repository) *git.Worktree {
	// Get the worktree
	worktree, err := repository.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	return worktree
}

func GetBranch(repository *git.Repository) string {
	head, err := repository.Head()
	if err != nil {
		log.Fatal(err)
	}

	return head.Name().Short()
}

func GetRemotes(repository *git.Repository) []*git.Remote {
	remotes, err := repository.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	return remotes
}

func GetRemoteNames(remotes []*git.Remote) []string {
	var names []string
	for _, remote := range remotes {
		names = append(names, remote.Config().Name)
	}

	return names
}
