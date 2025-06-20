package helper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

var configPath string = "/home/danielc/.config/swcli/config.json"

func TrimExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// GetConfigField gets the value for field from json file
func GetConfigField(fieldname string) string {
	// Read file content
	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("error reading file %s: %v", configPath, err))
	}

	// Validate JSON
	if !gjson.ValidBytes(jsonData) {
		panic(errors.New("invalid JSON format at config file"))
	}

	// Extract info
	return gjson.GetBytes(jsonData, fieldname).String()
}
