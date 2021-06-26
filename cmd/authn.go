package cmd

import "github.com/spf13/cobra"

var authnCmd = &cobra.Command{
	Use: "authn",
	Short: "Starts authentication service",
	Long: "Starts Customer service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(authnCmd)
}