package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/helper"
	"log"

	"github.com/go-git/go-git/v5"
)

func Commit(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	message, _ := cmd.Flags().GetString("message")
	amend, _ := cmd.Flags().GetBool("amend")
	name := helper.GetCurFuncName()

	if amend {
		log.Fatal("Not implemented yet...")
	}

	// Open the repository
	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "pre")

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	options := &git.CommitOptions{
		AllowEmptyCommits: false,
		All:               true,
	}

	// Commit changes with a commit message
	commit, err := worktree.Commit(message, options)
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "post")

	fmt.Printf("Committed: %s\n", commit)
}
