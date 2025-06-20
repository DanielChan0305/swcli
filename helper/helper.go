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

// jsonGetField gets the value of field from json file and returns as string
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
