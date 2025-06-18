package cmd

import (
	"errors"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/spf13/cobra"
)

var (
	starterTemplate string = "library/starterTemplate.cpp"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

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
		return nil
	},

	SilenceUsage: true,
}
