package template

import (
	"fmt"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// pathCmd outputs the path of the templates and copys to clipboard
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Outputs the path of the templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := helper.GetConfigField("templateFolder")
		fmt.Printf("✅ Path: %s\n", path)

		// copy to clipboard
		err := clipboard.WriteAll(string(path))

		if err != nil {
			return err
		}

		fmt.Printf("✅ Path copied to clipboard\n")
		return nil
	},
}
