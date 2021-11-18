package cli

import (
	"github.com/dunielpls/pano/pkg/server"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:     "pano",
		Short:   "Pano CLI",
		Version: "0.1",
	}
	verboseFlag bool
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "increase verbosity")

	cobra.EnableCommandSorting = true
	cobra.EnablePrefixMatching = false

	RootCmd.AddCommand(
		server.CLICommand(),
	)
}

func Execute() {
	// Executed by CLI entrypoint in `cmd`.
}
