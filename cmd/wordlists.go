package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var wordlistsCmd = &cobra.Command{
	Use:   "wordlists",
	Short: "Generate a install script for the most common wordlists",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GenerateWordlistsInstallScript()
	},
}

func init() {
	rootCmd.AddCommand(wordlistsCmd)
}
