package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Update objects and refs from remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		git.Fetch(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().BoolP("help", "h", false, "Show help for command")
	fetchCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
