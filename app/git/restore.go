package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"log"
)

func Restore(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	file := args[0]

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

	// Check the status of the file
	status, err := worktree.Status()
	if err != nil {
		log.Fatal(err)
	}

	if status.IsUntracked(file) {
		log.Println("File is not tracked.")
	}

	// Restore the file
	if err := worktree.Checkout(&git.CheckoutOptions{
		Hash:   plumbing.NewHash("HEAD^"),
		Create: false,
		Force:  true,
		Keep:   false,
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("File restored successfully!")
}
