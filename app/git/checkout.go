package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/helper"
	"log"
)

func Checkout(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	branch := args[0]
	name := helper.GetCurFuncName()
	force, _ := cmd.Flags().GetBool("force")

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

	// Checkout the desired branch
	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		Force:  force, // Use Force: true to discard local changes if needed
	})
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "post")

	fmt.Printf("Checked out to branch '%s'.\n", branch)
}
