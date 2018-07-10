package gopherenv

import (
	"log"
	"io/ioutil"
	"strings"
	"regexp"
)

var envContent = ""
var variables = make(map[string]string)

func getFile() string {
	if envContent == "" {
		file, err := ioutil.ReadFile("./.env")
		if err != nil {
			log.Println(err)
		}
		envContent = string(file)
	}
	return envContent
}

func setupValues() {
	getFile()
	lines := strings.Split(envContent, "\n")
	regex := regexp.MustCompile("^([a-zA-Z_-]*[0-9_-]*)[\\s]*=(.*)$")
	for _, line := range lines {
		if regex.MatchString(line) {
			variable := regex.FindAllStringSubmatch(line, -1)[0]
			variables[variable[1]] = variable[2]
		}
	}
}

func Env(key string) string {
	setupValues()
	if val, ok := variables[key]; ok {
		return val
	}
	return ""
}

