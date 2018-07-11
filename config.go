package gopherenv

import (
	"fmt"
	"os"
	"log"
)

var location = "."

func EnvLoc() string {
	return fmt.Sprintf("%s/.env", location)
}

// Set the path for the .env file of the project
func SetFolder(path string) error {
	if path[len(path) -1:] == "/" || path[len(path) -1:] == "\\" {
		path = path[:len(path) -1]
	}
	if _, err := os.Stat(fmt.Sprintf("%s/.env", path)); err != nil {
		log.Fatal(err)
		return err
	}
	location = path
	return nil
}