package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/DanielChan0305/swcli/helper"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFolder  string = "config"
	compileConfig string = "compile.json"
)

func init() {
	// init with viper
	cobra.OnInitialize(initConfig)

	// add commands
	rootCmd.AddCommand(compileCmd)

	// add flags
	compileCmd.Flags().Int("std", -1, "Select the std version for compilation")
}

/*
initConfig configs the compile function by loading default value of flags from .config file
*/
func initConfig() {
	viper.SetConfigName(helper.TrimExt(compileConfig))
	viper.AddConfigPath(configFolder)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("✅ Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("❌ Unable to read from config file, please check whether file exists")
		fmt.Println("❌", err)
	}
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
func buildExecutable(filename string, std int) error {
	statement := fmt.Sprintf("g++ %s -o %s -fdiagnostics-color", filename, helper.TrimExt(filename))
	// select std version
	statement += fmt.Sprintf(" -std=c++%d", std)

	// setup spinner and command
	s := spinner.New(spinner.CharSets[78], 100*time.Millisecond)
	s.Suffix = " " + statement + "\n"
	cmd := exec.Command("bash", "-c", statement)

	// activate spinner and start command
	s.Start()

	out, err := cmd.CombinedOutput()
	if err != nil {
		s.FinalMSG = fmt.Sprintf("❌ %s\n", err.Error())
	} else {
		s.FinalMSG = fmt.Sprintf("✅ %s\n", statement)
	}

	s.Stop()

	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
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
		// load the default value for std
		std := viper.GetInt("std")

		if tmp, err := cmd.Flags().GetInt("std"); err != nil {
			// error when reading std flag
			return fmt.Errorf("%w", err)
		} else if tmp != -1 {
			// overwrite default value with user value for std
			std = tmp
		}

		// start compiling
		err := buildExecutable(args[0], std)

		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmt.Println("✅ File compiled successfully")

		return nil
	},

	SilenceUsage: true,
}
