package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to the remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		git.Push(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.Flags().BoolP("help", "h", false, "Show help for command")
	pushCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
	pushCmd.Flags().StringP("user", "u", "", "Use the given <user>.")
	pushCmd.Flags().StringP("pass", "p", "", "Use the given <password>.")
}
