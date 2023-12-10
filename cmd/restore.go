package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore working tree files",
	Run: func(cmd *cobra.Command, args []string) {
		git.Restore(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().BoolP("help", "h", false, "Show help for command")
	restoreCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
