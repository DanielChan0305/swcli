package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
	//re := regexp.MustCompile(`^[A-Za-z0-9\\-\\._]+$`)

	//if !re.MatchString(filename) {
	//	return false, errors.New("filename contains invalid characters")
	//}

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

/*
buildExecutable builds an executable file fromo .cpp file

Returns the terminal output and errors and compile is unsucessfully
*/
func buildExecutable(filename string) (string, error) {
	var filenameWithoutExtension = strings.TrimSuffix(filename, filepath.Ext(filename))
	cmd := exec.Command("bash", "-c", fmt.Sprintf("g++ %s -o %s", filename, filenameWithoutExtension), "--color=force")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}

	return stdout.String(), nil
}

// compileCmd handles the compilation of executables from c++ source files
var compileCmd = &cobra.Command{
	Use:   "compile [filename]",
	Short: "Compiles the cpp file into executable",
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
		// start compiling
		msg, err := buildExecutable(args[0])

		if err != nil {
			fmt.Println("⚠️ Something went wrong !")
			fmt.Println()
			fmt.Fprint(os.Stderr, msg)

			return fmt.Errorf("%w", err)
		}

		fmt.Println("✅ File compiled successfully")

		return nil
	},

	SilenceUsage: true,
}
