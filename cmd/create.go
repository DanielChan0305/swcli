package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/spf13/cobra"
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

		// load from json file
		starterTemplatePath, err := helper.JsonGetFieldString(configTemplatePath, "starterTemplatePath")
		starterTemplatePath = filepath.Join(helper.GetExecPath(), starterTemplatePath)

		if err != nil {
			return err
		}

		// Execute the command
		cpCmd := exec.Command("bash", "-c", fmt.Sprintf("cp %s %s", starterTemplatePath, filename))
		output, err := cpCmd.CombinedOutput()

		if err != nil {
			fmt.Println(string(output))
			return err
		}

		fmt.Println("✅ File created successfully")

		return nil

	},

	SilenceUsage: true,
}
