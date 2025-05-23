package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFiles(url []string, rewrite bool) {
	count := len(url)

	for k, v := range url {
		path := LocalRepoData(v)
		fmt.Printf("\033[2K\r[%d / %d] ", k+1, count)

		if err := DownloadFile(v, path, rewrite); err != nil {
			fmt.Printf("[Error] DownloadFile - %w", err)
		}
	}
}

func DownloadFile(url, path string, rewrite bool) error {
	var err error

	if !rewrite {
		if _, err = os.Stat(path); err == nil {
			return nil
		}
	}

	if err = CreateDir(path); err != nil {
		fmt.Printf("[Error] DownloadFile - CreateDir - %s\n", path)
		return fmt.Errorf("CreateDir - %w", err)
	}

	out, err := os.Create(path)
	if err != nil {
		fmt.Printf("[Error] DownloadFile - os.Create - %s\n", path)
		return fmt.Errorf("os.Create - %w", err)
	}

	defer out.Close()
	fmt.Printf("Downloading: %s...", url)

	content, err := http.Get(url)
	if err != nil {
		fmt.Printf("[Error] DownloadFile - http.Get - %s\n", url)
		return fmt.Errorf("http.Get - %w", err)
	}

	defer content.Body.Close()

	if content.StatusCode == 404 {
		return fmt.Errorf("404 Not Found")
	}

	if _, err = io.Copy(out, content.Body); err != nil {
		fmt.Printf("[Error] DownloadFile - io.Copy - %s <= %s\n", url, path)
		return fmt.Errorf("io.Copy - %w", err)
	}

	return nil
}
