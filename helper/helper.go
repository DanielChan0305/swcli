package helper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

func TrimExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// JsonGetField gets the value of field from json file and returns as string
func JsonGetFieldString(filename string, fieldname string) (string, error) {
	// Read file content
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %v", filename, err)
	}

	// Validate JSON
	if !gjson.ValidBytes(jsonData) {
		return "", errors.New("invalid JSON format")
	}

	// Extract info
	return gjson.GetBytes(jsonData, fieldname).String(), nil
}

// GetExecPath returns the path of the executable, helping transform relative path to absolute path
// Panics if failed
func GetExecPath() string {
	exePath, err := os.Executable()
	if err != nil {
		panic("can't find valid path of executable")
	}

	exePath = filepath.Dir(exePath)

	return exePath
}
