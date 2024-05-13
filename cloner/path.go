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
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("[Error] Read file[%s] - %w\n", path, err)
		return "", err
	}

	return string(content), nil
}

func LocalRepoData(url string) string {
	return strings.Replace(url, repo, dataDir, -1)
}
