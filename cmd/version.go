package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd is the function which outputs the version number of Swcli
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version number of Swcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version of Swcli:  v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
