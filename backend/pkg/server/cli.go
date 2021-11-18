package server

import "github.com/spf13/cobra"

func CLICommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:          "server",
		Short:        "server",
		Hidden:       false,
		SilenceUsage: false,
		RunE: func(cmd *cobra.Command, args []string) error {
		},
	}

	return
}
