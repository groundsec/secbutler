package cmd

import (
	"fmt"
	"os"

	"github.com/groundsec/secbutler/pkg/logger"
	"github.com/groundsec/secbutler/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func completionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

var rootCmd = &cobra.Command{
	Use:   "secbutler",
	Short: "Essential utilities for pentester, bug-bounty hunters and security researchers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Read the help page using the -h flag to see the commands!")
	},
}

func init() {
	completion := completionCmd()
	completion.Hidden = true
	rootCmd.AddCommand(completion)
	logger.SetLevel(logrus.InfoLevel)
	utils.CheckAndCreateSecbutlerDir()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
