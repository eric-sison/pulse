package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// Iterate through parent directories to find the root directory containing a `go.mod` file.
func GetRootDir() (string, error) {

	currDir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	// Iterate through parent directories starting from the current
	// directory (`currDir`). It checks if a file named `go.mod` exists in the current directory using
	// `os.Stat`. If the file exists, it returns the current directory as the root directory. If the file
	// does not exist in the current directory, it moves up to the parent directory by setting `currDir`
	// to the parent directory using `filepath.Dir(currDir)`.
	for {
		if _, err := os.Stat(filepath.Join(currDir, "go.mod")); err == nil {
			return currDir, nil
		}

		rootDir := filepath.Dir(currDir)

		if rootDir == currDir {
			return "", fmt.Errorf("")
		}

		currDir = rootDir
	}
}
