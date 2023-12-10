package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/helper"
	"log"
	"os"
)

func Restore(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("path")
	files := args[0:]
	name := helper.GetCurFuncName()

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

	// Check the status of the file
	status, err := worktree.Status()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if status.IsUntracked(file) {
			log.Println("File is not tracked.")
		}

		if status.File(file).Worktree == git.Modified {

			// Restore the file to the state of the last commit
			hash := head.Hash()
			commit, err := repo.CommitObject(hash)
			if err != nil {
				log.Fatal("Commit:", err)
			}

			// Get the file content from the commit
			filePath, err := commit.File(file)
			if err != nil {
				log.Fatal("Content:", err)
			}

			// Read the file content from the commit
			content, err := filePath.Contents()
			if err != nil {
				log.Fatal(err)
			}

			perm := filePath.Mode
			name := fmt.Sprintf("%s/%s", dir, file)

			// Reset the file permissions
			os.Chmod(name, os.FileMode(perm))

			// Write the file content to the working directory
			os.WriteFile(name, []byte(content), os.FileMode(perm))

			fmt.Printf("Restored: %s\n", file)
		}
	}
}
