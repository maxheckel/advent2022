package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInputLines(path string) ([]string, error) {
	if path[0] != '/' {
		dir, _ := os.Getwd()
		path = fmt.Sprintf("%s/%s", dir, path)
	}
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(dat), "\n"), nil
}

func ReadWholeFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
