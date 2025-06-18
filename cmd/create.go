package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	starterTemplate string = "library/starterTemplate.cpp"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func isFilenameValid(filename string) (bool, error) {
	if filename == "" {
		return false, errors.New("filename can't be empty")
	}

	// check if filename matches the required regex (allowing relative and absolute paths, alphanumerics, underscores, hyphens, and dots)
	matched, err := regexp.MatchString(`^([a-zA-Z0-9_\-./]+)$`, filename)
	if err != nil {
		return false, fmt.Errorf("error matching filename regex: %w", err)
	}
	if !matched {
		return false, errors.New("filename contains invalid characters or format")
	}

	// check whether it is a cpp file
	ext := filepath.Ext(filename)

	if ext != ".cpp" {
		return false, errors.New("only .cpp files can be created")
	}

	return true, nil
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

		if _, err := isFilenameValid(args[0]); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},

	SilenceUsage: true,
}
