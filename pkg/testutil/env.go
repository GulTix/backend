package testutil

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	_, fileName, _, _ := runtime.Caller(0)

	dir := filepath.Dir(fileName)

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// Found the project root
			return godotenv.Load(filepath.Join(dir, ".env"))
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached the file system root without finding go.mod
			return os.ErrNotExist
		}
		dir = parent
	}
}
