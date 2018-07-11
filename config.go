package gopherenv

import (
	"fmt"
	"os"
	"log"
)

var folder = "./"

func EnvLoc() string {
	return fmt.Sprintf("%s.env", folder)
}

// Set the path for the .env file of the project
func SetFolder(path string) error {
	if path[len(path) -1:] == "/" || path[len(path) -1:] == "\\" {
		path = path[:len(path) -1]
	}
	if _, err := os.Stat(fmt.Sprintf("%s.env", path)); err != nil {
		log.Fatal(err)
	}
	folder = path
}