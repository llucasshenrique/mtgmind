package database

import (
	"fmt"
	"os"
)

func ClearDatabase(path string) {
	if CheckIfFileExists(path) {
		err := os.Remove(path)
		if err != nil {
			fmt.Printf("%v", err)
		}
	} else {
		fmt.Println("Database does not exist")
	}
}
