package template

import "github.com/spf13/cobra"

// listCmd lists all custom templates
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all custom templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
