package database

import (
	"errors"
	"os"
)

func CheckIfFileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
