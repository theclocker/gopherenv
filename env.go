package gopherenv

import (
	"log"
	"io/ioutil"
	"strings"
	"regexp"
	"sync"
	"fmt"
)

var envContent = ""
var variables = make(map[string]string)
var (
	once sync.Once
)
func getFile() string {
	once.Do(func() {
		file, err := ioutil.ReadFile(EnvLoc())
		if err != nil {
			log.Println(err)
		}
		envContent = string(file)
	})
	return envContent
}

func setupValues() {
	getFile()
	lines := strings.Split(envContent, "\n")
	regex := regexp.MustCompile("^([a-zA-Z]+[a-zA-Z0-9_-]*)[\\s]*=(.*)$")
	for _, line := range lines {
		if regex.MatchString(line) {
			variable := regex.FindAllStringSubmatch(line, -1)[0]
			variables[variable[1]] = variable[2]
		}
	}
}

func validateKey(key string) bool {
	regex := regexp.MustCompile("^[a-zA-Z]+[a-zA-Z0-9_-]*$")
	return regex.MatchString(key)
}

func Env(key string) string {
	if validateKey(key) {
		setupValues()
		if val, ok := variables[key]; ok {
			return val
		} else {
			log.Fatal(fmt.Sprintf("Key: %s does not exist", key))
		}
	}
	log.Fatal(fmt.Sprintf("Key: %s is not valid", key))
	return ""
}

