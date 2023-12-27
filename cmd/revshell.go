package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var revshellCmd = &cobra.Command{
	Use:   "revshell",
	Short: "Obtain the command for a reverse shell",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GetReverseShell()
	},
}

func init() {
	rootCmd.AddCommand(revshellCmd)
}
