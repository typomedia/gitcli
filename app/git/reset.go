package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var options = &git.ResetOptions{}

func Reset(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	hard, _ := cmd.Flags().GetBool("hard")

	// Open the repository
	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	if hard {
		options = &git.ResetOptions{
			Mode:   git.HardReset,
			Commit: plumbing.NewHash("HEAD^"),
		}
	} else {
		options = &git.ResetOptions{
			Mode:   git.SoftReset,
			Commit: plumbing.NewHash("HEAD^"),
		}
	}

	// Reset the working directory to the state of the last commit
	err = worktree.Reset(options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Changes stashed successfully.")
}
