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

	// Retrieve the HEAD reference
	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	err = repo.Fetch(&git.FetchOptions{})
	if err != nil {
		//log.Fatal(err)
	}

	hash := head.Hash()
	ref := plumbing.NewHashReference(plumbing.ReferenceName(branch), hash)

	err = repo.Storer.RemoveReference(ref.Name())
	if err != nil {
		log.Fatal(err)
	}

	// Checkout the desired branch
	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branch),
		Create: true,
		Force:  force,
	})
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "post")

	fmt.Printf("Checked out to branch '%s'.\n", branch)
}
