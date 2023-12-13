package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
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

	err = repo.Fetch(&git.FetchOptions{})
	if err != nil {
		//log.Fatal(err)
	}

	branches, _ := repo.Branches()
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == branch {
			// delete target branch before
			err := repo.Storer.RemoveReference(ref.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	//// Retrieve the HEAD reference
	//head, err := repo.Head()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// Get the commit object
	//hash := head.Hash()
	//ref := plumbing.NewHashReference(plumbing.ReferenceName(branch), hash)
	//
	//// delete target branch before
	//fmt.Println(ref.Name())
	//err = repo.Storer.RemoveReference(ref.Name())
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Carefull!!
	//if force {
	//	err = repo.Storer.RemoveReference(plumbing.NewRemoteReferenceName("origin", branch))
	//	if err != nil {
	//		log.Fatal("RemoveReference", err)
	//	}
	//}

	remotes, err := repo.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(remotes)

	hash := plumbing.NewHash("")
	for _, remote := range remotes {
		// Fetch the branches for each remote
		refs, err := remote.List(&git.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		// Find the specific branch
		for _, ref := range refs {
			if ref.Name().IsBranch() {
				if ref.Name().Short() == branch {
					hash = ref.Hash()
				}

			}
		}
	}

	fmt.Println(hash)

	// Checkout the desired branch
	err = worktree.Checkout(&git.CheckoutOptions{
		//Branch: plumbing.NewRemoteReferenceName("origin", branch),
		Hash: plumbing.NewHash(hash.String()),
		//Hash:   hash,
		Branch: plumbing.NewBranchReferenceName(branch),
		Create: true,
		Force:  force,
	})
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := repo.Config()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Branches["main"])

	// Set branch configuration
	remoteName := "origin"
	branchConfig := &config.Branch{
		Name:   branch,
		Remote: remoteName,
		Merge:  plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	}
	cfg.Branches[branch] = branchConfig

	fmt.Println(branchConfig)

	// Save the updated configuration
	err = repo.Storer.SetConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "post")

	fmt.Printf("Checked out to branch '%s'.\n", branch)
}
