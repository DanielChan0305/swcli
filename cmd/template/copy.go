package template

import (
	"errors"
	"fmt"
	"os"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// copyCmd copys the desired template into clipboard
var copyCmd = &cobra.Command{
	Use:   "copy [template name]",
	Short: "Copys the desired template into clipboard",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("template name can't be empty")
		}

		templateName := args[0]
		if _, err := helper.IsFilenameValid(templateName); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		templateName := args[0]
		templateFolder, err := helper.JsonGetFieldString(configTemplatePath, "templateFolder")

		if err != nil {
			return err
		}

		templatePath := templateFolder + "/" + templateName + ".h"
		//fmt.Println(templatePath)
		if !helper.IsFileExist(templatePath) {
			return fmt.Errorf("can't find template: %s", templateName)
		}

		fmt.Printf("✅ Found template: %s\n", helper.TrimExt(templatePath))
		templateContent, err := os.ReadFile(templatePath)

		if err != nil {
			return err
		}

		err = clipboard.WriteAll(string(templateContent))
		if err != nil {
			return err
		}

		fmt.Printf("✅ Copied %s into clipboard\n", helper.TrimExt(templatePath))

		return nil
	},

	SilenceUsage: true,
}
