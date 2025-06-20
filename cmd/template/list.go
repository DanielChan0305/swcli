package template

import (
	"fmt"
	"os"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/spf13/cobra"
)

var (
	// the path for the configuration file of template function
	configTemplatePath = "config/template.json"
)

// listCmd lists all custom templates
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all custom templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		// the stored value in the configuration file
		// indicating the path of the folder storing the templates
		templateFolder, err := helper.JsonGetFieldString(configTemplatePath, "templateFolder")

		if err != nil {
			return err
		}

		templates, err := os.ReadDir(templateFolder)

		if err != nil {
			return err
		}

		for ix, template := range templates {
			fmt.Printf("%d -- %s\n", ix, helper.TrimExt(template.Name()))
		}

		return nil
	},
}
