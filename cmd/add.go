package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file contents to the index",
	Run: func(cmd *cobra.Command, args []string) {
		git.Add(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("help", "h", false, "Show help for command")
	addCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
