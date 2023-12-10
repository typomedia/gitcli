package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/helper"
	"log"

	"github.com/go-git/go-git/v5"
)

func Pull(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	name := helper.GetCurFuncName()

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

	// Pull changes from the default remote branch (usually origin/master)
	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatal(err)
	}

	Hook(dir, name, "post")

	fmt.Println("Pull operation completed.")
}
