package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is inits the root command of swcli and is the gateway to other functions
var rootCmd = &cobra.Command{
	Use:   "swcli",
	Short: "Swcli helps you compile your C++ code more efficiently and enables easy importing of template code.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
