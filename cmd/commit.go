package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record changes to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		git.Commit(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().BoolP("help", "h", false, "Show help for command")
	commitCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
	commitCmd.Flags().StringP("message", "m", "", "Use the given <msg> as the commit message.")
	commitCmd.Flags().BoolP("amend", "a", false, "Amend the tip of the current branch using the same log message as <commit>.")
}
