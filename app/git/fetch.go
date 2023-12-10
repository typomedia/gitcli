package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"log"
)

func Fetch(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	err = repo.Fetch(&git.FetchOptions{})
	if err != nil {
		log.Fatal(err)
	}
}
