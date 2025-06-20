package cmd

import (
	"errors"
	"fmt"
	"os/exec"
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

/*
viperConfg loads the default value of flags from .config file
*/
func viperConfig() {
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
buildExecutable builds an executable file fromo .cpp file

Pipes the terminal output and errors, return error is compile is unsuccessful
*/
func buildExecutable(filename string, std int) error {
	statement := fmt.Sprintf("g++ %s -o %s -fdiagnostics-color", filename, helper.TrimExt(filename))
	// select std version
	statement += fmt.Sprintf(" -std=c++%d ", std)
	// add default compilation flags
	statement += "-O2 -Wall -Wextra -Wshadow -fsanitize=address,undefined"

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
	Use:   "compile [filename (include extension)]",
	Short: "Compiles the cpp file into executable",
	Args: func(cmd *cobra.Command, args []string) error {
		// validates the filename
		if len(args) < 1 {
			return errors.New("filename can't be empty")
		}

		filename := args[0]
		if _, err := helper.IsFilenameValid(filename); err != nil {
			return err
		}

		// check whether file exists
		if !helper.IsFileExist(filename) {
			return errors.New("file doesn't exist")
		}

		// check whether it is a cpp file
		if !helper.IsCpp(filename) {
			return errors.New("only .cpp file can be compiled")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		viperConfig()

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
