package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Generate a install script for the most common cybersecurity tools",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GenerateToolsInstallScript()
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)
}
