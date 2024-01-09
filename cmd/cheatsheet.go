package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var cheatsheetCmd = &cobra.Command{
	Use:   "cheatsheet",
	Short: "Read common cheatsheets & payloads",
	Run: func(cmd *cobra.Command, args []string) {
		runners.ShowCheatsheet()
	},
}

func init() {
	rootCmd.AddCommand(cheatsheetCmd)
}
