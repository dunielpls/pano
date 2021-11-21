package cli

import (
	"fmt"
	"os"

	"github.com/dunielpls/pano/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:     "pano",
		Short:   "Pano CLI",
		Version: "0.1",
	}
	srv        server.Server
	colorsFlag bool
	configFile string

	// TODO: Log levels instead of verbosity.
	verboseFlag bool
)

func init() {
	// CLI initialization.
	cobra.EnableCommandSorting = true
	cobra.EnablePrefixMatching = false

	// Run `initHook` before each command is executed.
	cobra.OnInitialize(initHook)

	// Root-level Cobra flags.
	rootCmd.PersistentFlags().BoolVarP(&colorsFlag, "no-color", "n", false, "disable colors")
	// TODO: Log levels instead of verbosity.
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "/etc/pano.conf", "path to configuration file")

	// Disable the default completion` command.
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(&cobra.Command{
		Use:          "version",
		Short:        "Show version and exit",
		Hidden:       false,
		SilenceUsage: false,
		Run: func(cmd *cobra.Command, args []string) {
			// Entrypoint for `pano version`.
			fmt.Print(srv.CLIVersion())
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:          "status",
		Short:        "Show status and exit",
		Hidden:       false,
		SilenceUsage: false,
		Run: func(cmd *cobra.Command, args []string) {
			// Entrypoint for `pano status`.
			fmt.Print(srv.CLIStatus())
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:          "start",
		Short:        "Start the server",
		Hidden:       false,
		SilenceUsage: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Entrypoint for `pano start`.
			return srv.CLIStart()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:          "stop",
		Short:        "Stop the server",
		Hidden:       false,
		SilenceUsage: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Entrypoint for `pano stop`.
			return srv.CLIStop()
		},
	})
}

func initHook() {
	// Basic configuration.
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("pano")
	viper.AutomaticEnv()

	// Defaults
	viper.SetDefault("server.bind", "0.0.0.0")
	viper.SetDefault("server.port", 8081)
	// TODO: Implement.
	viper.SetDefault("server.trusted_proxies", "")
	viper.SetDefault("server.routes.view", "/view")
	viper.SetDefault("server.routes.edit", "/edit")
	viper.SetDefault("server.routes.api_prefix", "/api/v1")
	viper.SetDefault("zabbix.url", "https://zabbix/zabbix/api_jsonrpc.php")
	viper.SetDefault("zabbix.username", "sa_pano")
	viper.SetDefault("zabbix.password", "sa_pano")

	// Read config file.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Error parsing configuration file: %v\n", err)
		}
	}

	srv = server.New()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error running rootCmd.Execute(): %v\n", err)
		os.Exit(1)
	}
}
