package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var (
	configTemplatePath string = "config/template.json"
)

// createCmd creates a new file based on the starter template
var createCmd = &cobra.Command{
	Use:   "create [filename]",
	Short: "Create the cpp file with starter template",
	Args: func(cmd *cobra.Command, args []string) error {
		// validates the filename
		if len(args) < 1 {
			return errors.New("filename can't be empty")
		}

		filename := args[0]
		if !helper.IsCpp(filename) {
			return errors.New("only .cpp files can be created")
		}

		if _, err := helper.IsFilenameValid(filename); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		// Read file content
		jsonData, err := os.ReadFile(configTemplatePath)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", configTemplatePath, err)
		}

		// Validate JSON
		if !gjson.ValidBytes(jsonData) {
			return errors.New("invalid JSON format")
		}

		// Extract info
		starterTemplatePath := gjson.GetBytes(jsonData, "starterTemplatePath").String()

		// Execute the command
		cpCmd := exec.Command("bash", "-c", fmt.Sprintf("cp %s %s", starterTemplatePath, filename))
		err = cpCmd.Run()

		if err != nil {
			return err
		}

		fmt.Println("✅ File created successfully")

		return nil

	},

	SilenceUsage: true,
}
