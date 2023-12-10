package cmd

import (
	"github.com/spf13/cobra"
	"github.com/typomedia/gitcli/app/git"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Switch branches or restore working tree files",
	Run: func(cmd *cobra.Command, args []string) {
		git.Checkout(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().BoolP("help", "h", false, "Show help for command")
	checkoutCmd.Flags().StringP("path", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
}
