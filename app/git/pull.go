package git

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git/common"
	"github.com/typomedia/gitcli/app/helper"
	"log"

	"github.com/go-git/go-git/v5"
)

func Pull(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	name := helper.GetCurFuncName()

	// Open the repository
	repo := common.GetRepo(dir)
	branch := common.GetBranch(repo)

	Hook(dir, name, "pre")

	// Get the worktree
	worktree := common.GetWorktree(repo)

	// Get the remotes
	remotes := common.GetRemotes(repo)
	names := common.GetRemoteNames(remotes)

	// Pull changes from the remote branch
	err := worktree.Pull(&git.PullOptions{
		RemoteName:    names[0],
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatal(err)
	}

	Hook(dir, name, "post")

	fmt.Println("Already up to date.")
}
