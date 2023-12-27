package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var payloadsCmd = &cobra.Command{
	Use:   "payloads",
	Short: "Obtain and serve common payloads",
	Run: func(cmd *cobra.Command, args []string) {
		runners.ServePayloads()
	},
}

func init() {
	rootCmd.AddCommand(payloadsCmd)
}
