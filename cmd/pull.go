package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls changes from the remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		git.Pull(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().BoolP("help", "h", false, "Show help for command")
	pullCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
