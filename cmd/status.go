package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		git.Status(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().BoolP("help", "h", false, "Show help for command")
	statusCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
