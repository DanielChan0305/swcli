package template

import "github.com/spf13/cobra"

// getCmd copys the desired template into clipboard
var getCmd = &cobra.Command{
	Use:   "get [template name]",
	Short: "Copys the desired template into clipboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
