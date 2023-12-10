package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset current HEAD to the specified state",
	Run: func(cmd *cobra.Command, args []string) {
		git.Reset(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.Flags().BoolP("help", "h", false, "Show help for command")
	resetCmd.Flags().BoolP("hard", "", false, "Reset the index and working tree")
	resetCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
