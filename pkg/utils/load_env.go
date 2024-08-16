package utils

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Load environment variables from a .env file located in the root directory of the project.
func LoadEnv() {
	rootDir, err := GetRootDir()

	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(rootDir, ".env"))

	if err != nil {
		log.Fatal(err)
	}

}
