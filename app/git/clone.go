package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"path"
)

func Clone(cmd *cobra.Command, args []string) {
	repo := args[0]
	dir := args[1]

	// Parse the URL
	u, err := url.Parse(repo)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Extract the last part of the path
	clonePath := path.Base(u.Path)
	if dir != "." {
		clonePath = dir
	}

	// Clone the repository
	_, err = git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout, // Optional: To display cloning progress, you can set Progress to os.Stdout
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository cloned successfully.")
}
