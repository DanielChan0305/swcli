package template

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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
		// load templateName and templateFolder from config files
		templateName := args[0]
		templateFolder := helper.GetConfigField("templateFolder")

		// check whether path is valid
		templatePath := filepath.Join(templateFolder, templateName+".h")
		if !helper.IsFileExist(templatePath) {
			return fmt.Errorf("can't find template: %s", templateName)
		}

		// read from path
		templateContent, err := os.ReadFile(templatePath)

		if err != nil {
			return err
		}

		// copy to clipboard
		err = clipboard.WriteAll(string(templateContent))

		if err != nil {
			return err
		}

		fmt.Printf("✅ Copied from template into clipboard: %s\n", helper.TrimExt(templatePath))

		return nil
	},

	SilenceUsage: true,
}
