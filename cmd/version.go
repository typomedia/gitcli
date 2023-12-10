package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of gitcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gitcli version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
