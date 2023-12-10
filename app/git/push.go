package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/helper"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Push(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	user, _ := cmd.Flags().GetString("user")
	token, _ := cmd.Flags().GetString("pass")
	name := helper.GetCurFuncName()

	// Open the repository
	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "pre")

	// Fetch all remotes configured for the repository
	remotes, err := repo.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over remotes and push changes using the first remote found (change this logic as needed)
	var selectedRemote *git.Remote
	for _, remote := range remotes {
		// Choose the remote based on your criteria (e.g., fetch URL, name, etc.)
		// For example, you might choose the first remote found
		selectedRemote = remote
		break
	}

	if selectedRemote == nil {
		log.Fatal("No remotes found")
	}

	// Set up authentication with username and token
	auth := &http.BasicAuth{
		Username: user,
		Password: token,
	}

	// Push commits to the remote repository with authentication
	err = selectedRemote.Push(&git.PushOptions{
		Auth: auth,
	})
	if err != nil {
		log.Fatal(err)
	}

	Hook(dir, name, "post")

	fmt.Println("Pushed commits to remote repository.")
}
