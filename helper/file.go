package helper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// isCpp checks whether the file is a c++ file
func IsCpp(filename string) bool {
	return (filepath.Ext(filename) == ".cpp")
}

// isFileExist checks whether a file (excluding directorys) of the given name exists
func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

// isFilenameValid checks whether the filename is valid and contains invalid characters
func IsFilenameValid(filename string) (bool, error) {
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

	return true, nil
}
