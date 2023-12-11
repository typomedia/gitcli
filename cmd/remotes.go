package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var remotesCmd = &cobra.Command{
	Use:   "remotes",
	Short: "List",
	Run: func(cmd *cobra.Command, args []string) {
		git.Remotes(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(remotesCmd)
	remotesCmd.Flags().BoolP("help", "h", false, "Show help for command")
	remotesCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
