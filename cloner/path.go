package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateDir(path string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func GetContentFromFile(path string) (string, error) {
	if _, err := os.Stat(path); err != nil {
		return "", nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("[Error] Read file[%s] - %s\n", path, err.Error())
		return "", err
	}

	return string(content), nil
}

func LocalRepoData(url string) string {
	return strings.Replace(url, repo, dataDir, -1)
}
