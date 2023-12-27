package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Obtain the command to start a reverse shell listener",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GetReverseShellListener()
	},
}

func init() {
	rootCmd.AddCommand(listenerCmd)
}
