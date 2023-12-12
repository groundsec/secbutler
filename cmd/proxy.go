package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Obtain a random proxy from FreeProxy",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GetRandomProxy()
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
