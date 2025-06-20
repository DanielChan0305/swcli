package template

import (
	"fmt"
	"path/filepath"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/spf13/cobra"
)

// listCmd lists all custom templates
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all custom templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		// the stored value in the configuration file
		// indicating the path of the folder storing the templates
		templateFolder, err := helper.JsonGetFieldString(configTemplatePath, "templateFolder")
		templateFolder = filepath.Join(helper.GetExecPath(), templateFolder)

		if err != nil {
			return err
		}

		// get files with .h extension
		templates, err := filepath.Glob(templateFolder + "/*.h")

		if err != nil {
			return err
		}

		// displaying the template files
		for ix, template := range templates {
			fmt.Printf("%d - %s\n", ix, helper.TrimExt(filepath.Base(template)))
		}

		return nil
	},
}
