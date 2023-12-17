package cmd

import (
	"github.com/groundsec/secbutler/pkg/runners"
	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Obtain a random proxy from FreeProxy",
	Run: func(cmd *cobra.Command, args []string) {
		untestedFlag, _ := cmd.Flags().GetBool("untested")
		randomFlag, _ := cmd.Flags().GetBool("random")
		httpsFlag, _ := cmd.Flags().GetBool("https")
		eliteFlag, _ := cmd.Flags().GetBool("elite")
		anonFlag, _ := cmd.Flags().GetBool("anon")
		googleFlag, _ := cmd.Flags().GetBool("google")
		countriesFlag, _ := cmd.Flags().GetString("countries")
		runners.GetProxy(untestedFlag, randomFlag, httpsFlag, eliteFlag, anonFlag, googleFlag, countriesFlag)
	},
}

func init() {
	proxyCmd.Flags().Bool("untested", false, "Tested proxy")
	proxyCmd.Flags().Bool("random", false, "Random proxy")
	proxyCmd.Flags().Bool("https", false, "HTTPS proxy")
	proxyCmd.Flags().Bool("elite", false, "Elite proxy")
	proxyCmd.Flags().Bool("anon", false, "Anonymous proxy")
	proxyCmd.Flags().Bool("google", false, "Proxies with 'Google' set to 'yes'")
	proxyCmd.Flags().String("countries", "", "The countries of the proxy (comma separated values)")

	rootCmd.AddCommand(proxyCmd)
}
