package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is inits the root command of swcli and is the gateway to other functions
var rootCmd = &cobra.Command{
	Use:   "swcli",
	Short: "Swcli helps you compile your C++ code more efficiently and enables easy importing of template code.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Show help when no subcommands or arguments are provided
		return cmd.Help()
	},
}

func Execute() {
	rootCmd.SetErrPrefix("❌")
	if err := rootCmd.Execute(); err != nil {
		//fmt.Fprintln(os.Stderr, "Error [root.go] :", err)
		os.Exit(1)
	}
}
