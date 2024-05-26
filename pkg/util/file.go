// Package util contains utilities.
package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetFile is a function to get a file.
func GetFile(filename string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dirs := strings.Split(dir, "/")

	for i := len(dirs); i > 0; i-- {
		dir := "/"
		for j := 0; j < i; j++ {
			dir = filepath.Join(dir, dirs[j])
		}

		file := filepath.Join(dir, filename)
		fileInfo, err := os.Stat(file)
		if os.IsNotExist(err) {
			continue
		}
		if !fileInfo.IsDir() {
			return file, nil
		}
	}

	return "", fmt.Errorf("%s not found", filename)
}
