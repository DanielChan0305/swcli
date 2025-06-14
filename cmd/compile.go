package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(compileCmd)
}

/*
isFilenameValid is a helper function, which validates the file name

Returns false if filename contains invalid characters or file doesn't exists

Otherwise, Returns true
*/
func isFilenameValid(filename string) (bool, error) {
	if filename == "" {
		return false, errors.New("filename can't be empty")
	}

	// check whether filename contains invalid characters
	re := regexp.MustCompile(`^[A-Za-z0-9\\-\\._]+$`)

	if !re.MatchString(filename) {
		return false, errors.New("filename contains invalid characters")
	}

	// check whether file exists
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, errors.New("file doesn't exist")
	}

	if info.IsDir() {
		return false, errors.New("please enter filename not directory name")
	}

	// check whether it is a cpp file
	ext := filepath.Ext(filename)

	if ext != ".cpp" {
		return false, errors.New("only .cpp files can be compiled")
	}

	return true, nil
}

// compileCmd handles the compilation of executables from c++ source files
var compileCmd = &cobra.Command{
	Use:   "compile [FILENAME]",
	Short: "Compiles the cpp file into executable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("filename can't be empty")
		}

		if _, err := isFilenameValid(args[0]); err != nil {
			return fmt.Errorf("%w", err)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// start compiling



		fmt.Println("File compiled successfully")
		return nil
	},
}
