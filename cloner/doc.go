package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateDoc(track OSTrack) {
	stub := track.Stub()
	urlRepo := track.UrlRepo()

	content, err := GetContentFromFile(stub)
	if err != nil {
		fmt.Printf("[Error] Error read stub - %v", err)
		return
	}

	name := fmt.Sprintf("%s (%s)", strings.ToTitle(track.OS), strings.ToTitle(track.Version))

	if track.Version == "" {
		name = strings.ToTitle(track.OS)
	}

	content = strings.ReplaceAll(content, "<NAME>", name)

	if len(urlRepo) > 1 {
		content = strings.ReplaceAll(content, "<SIGNER>", urlRepo[1])
		content = strings.ReplaceAll(content, "<REPO>", urlRepo[0])
	} else {
		return
	}

	docName := track.Version
	if docName == "" {
		docName = track.OS
	}

	docPath := fmt.Sprintf("%s/%s/%s/%s.md", docDir, track.Channel, track.OS, docName)
	if err := CreateDir(docPath); err != nil {
		fmt.Printf("[Error] Error write doc - %v", err)
		return
	}

	if err := os.WriteFile(docPath, []byte(content), 0644); err != nil {
		fmt.Printf("[Error] Error write stub - %v", err)
		return
	}

	if err := ReplaceSourceRepoToMirror(docPath); err != nil {
		fmt.Printf("[Error] Error replace repo to mirror in stub - %v", err)
		return
	}
}
