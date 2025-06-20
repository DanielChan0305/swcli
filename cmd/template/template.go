package template

import (
	"github.com/spf13/cobra"
)

// templateCmd is the function which supports the use and import of custom templates
var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Supports the use and import of custom templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	TemplateCmd.AddCommand(getCmd)
	TemplateCmd.AddCommand(listCmd)
}
