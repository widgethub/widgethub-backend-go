package cmd

import (
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use: "all",
	Short: "Starts all services",
	Long: "Starts all sevcies",
	RunE: func(cmd *cobra.Command, args []string) error {
		go authnCmd.RunE(authnCmd, args)
		return serverCmd.RunE(serverCmd, args)
	},
}

func init() {
	RootCmd.AddCommand(allCmd)
}