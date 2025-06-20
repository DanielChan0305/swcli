package template

import (
	"path/filepath"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/spf13/cobra"
)

var (
	// the path for the configuration file of template function
	configTemplatePath = "config/template.json"
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
	configTemplatePath = filepath.Join(helper.GetExecPath(), configTemplatePath)

	TemplateCmd.AddCommand(copyCmd)
	TemplateCmd.AddCommand(listCmd)
}
