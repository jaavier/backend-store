package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadEnv(path ...string) error {
	var finalPath string
	if len(path) == 0 {
		finalPath = "./.env"
	} else {
		finalPath = path[0]
	}
	fileContent, err := readFile(finalPath)
	if err != nil {
		return err
	}
	err = DefineEnv(fileContent)
	if err != nil {
		return err
	}
	return nil
}

func DefineEnv(content string) error {
	var byLines []string = strings.Split(content, "\n")
	if len(byLines) == 0 {
		return fmt.Errorf("cannot define environment variables")
	}
	for _, line := range byLines {
		var splitLine []string = strings.Split(line, "=")
		var name string = splitLine[0]
		var value string = strings.Join(splitLine[1:], "")
		os.Setenv(name, value)
	}
	return nil
}

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading %s", path)
	}
	return string(content), nil
}
