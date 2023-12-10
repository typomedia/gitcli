package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/helper"
	"log"

	"github.com/go-git/go-git/v5"
)

func Add(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	files := args[0:]
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

	for _, file := range files {
		// Add the file to the staging area
		_, err = worktree.Add(file)
		if err != nil {
			log.Fatal("failed to stage changes: %s", err)
		}
	}

	Hook(dir, name, "post")

	if len(files) > 1 {
		fmt.Println("Changes staged for commit.")
	} else {
		fmt.Println("Change staged for commit.")
	}
}
