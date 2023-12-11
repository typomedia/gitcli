package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/helper"
	"log"
)

func Remotes(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	name := helper.GetCurFuncName()

	// Open the repository
	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "pre")

	err = repo.Fetch(&git.FetchOptions{})
	if err != nil {
		//log.Fatal(err)
	}

	remotes, err := repo.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	branches := Refs{}

	// Loop through remotes
	for _, remote := range remotes {
		// Fetch the branches for each remote
		refs, err := remote.List(&git.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		// Find the specific branch
		for _, ref := range refs {
			if ref.Name().IsBranch() {
				// Only show remote branches
				branches.Branches = append(branches.Branches, Branch{
					Name:     ref.Name().Short(),
					Revision: ref.Hash().String(),
				})

			}
		}
	}

	fmt.Println(branches)
}

type Refs struct {
	Branches []Branch
}

type Branch struct {
	Name     string
	Revision string
}
