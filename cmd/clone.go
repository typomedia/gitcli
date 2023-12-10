package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository into a new directory",
	Run: func(cmd *cobra.Command, args []string) {
		git.Clone(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().BoolP("help", "h", false, "Show help for command")
}
