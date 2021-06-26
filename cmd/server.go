package cmd

import (
	"net"
	"widgethub/v2/services/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Starts provider service",
	Long: "Starts provider service",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := server.NewServer(net.JoinHostPort("0.0.0.0", "8080"))
		return server.Run()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}