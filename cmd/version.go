package cmd

import (
	"fmt"

	"github.com/groundsec/secbutler/pkg/utils"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%sCurrent version%s: v%s\n", utils.ANSICodes["Bold"], utils.ANSICodes["Reset"], utils.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
