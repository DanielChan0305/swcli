package helper

import (
	"path/filepath"
	"strings"
)

func TrimExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
