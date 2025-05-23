package main

import (
	"errors"
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
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Println("Error read content: %w", err)
		}

		return "", nil
	}

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
