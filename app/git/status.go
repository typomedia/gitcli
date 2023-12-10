package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"log"
)

func Status(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")

	// Open the repository
	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the HEAD reference
	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	// Get the branch name from the HEAD reference
	branchName := head.Name().Short()

	fmt.Printf("On branch: %s\n", branchName)

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	// Get the status of the worktree
	status, err := worktree.Status()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(status)

	// print revision number
	headHash := head.Hash()
	fmt.Println(headHash)
}
