package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReplaceSourceRepoToMirror(url string) error {
	path := LocalRepoData(url)

	input, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	lines := strings.Split(string(input), "\n")
	for k, v := range lines {
		if strings.Contains(v, repo) {
			lines[k] = strings.Replace(lines[k], repo, mirror, -1)
		}
		if strings.Contains(v, "tailscale-archive-keyring") {
			lines[k] = strings.Replace(lines[k], "tailscale-archive-keyring", "tailscale-mirror-archive-keyring", -1)
		}
	}

	output := strings.Join(lines, "\n")

	return os.WriteFile(path, []byte(output), 0644)
}

func RepoDist(t OSTrack) (string, string) {
	path := fmt.Sprintf("%s/%s/%s", repo, t.Channel, t.OS)

	if t.Version == "" {
		return path, fmt.Sprintf("%s/dists", path)
	}

	return path, fmt.Sprintf("%s/dists/%s", path, t.Version)
}

func RepoMD(t OSTrack, arch string) (string, string) {
	repoVersion := t.Version
	dist := fmt.Sprintf("%s/%s/%s/%s/%s", repo, t.Channel, t.OS, repoVersion, arch)

	if repoVersion == "" {
		dist = fmt.Sprintf("%s/%s/%s/%s", repo, t.Channel, t.OS, arch)
	}

	md := fmt.Sprintf("%s/repodata/repomd.xml", dist)

	return md, dist
}

func RepoData(t OSTrack, arch, repoMd string) (string, []string, error) {
	path := LocalRepoData(repoMd)
	content, err := GetContentFromFile(path)
	if err != nil {
		return "", nil, fmt.Errorf("GetContentFromFile - %w", err)
	}

	if content == "" {
		return "", nil, fmt.Errorf("no repo data")
	}

	primary, files, err := ParseRepoData(content, 1)
	if err != nil {
		return "", nil, fmt.Errorf("ParseRepoData: %w", err)
	}

	repoVersion := t.Version
	primaryPath := fmt.Sprintf("%s/%s/%s/%s/%s/%s", repo, t.Channel, t.OS, repoVersion, arch, primary)

	if repoVersion == "" {
		primaryPath = fmt.Sprintf("%s/%s/%s/%s/%s", repo, t.Channel, t.OS, arch, primary)
	}

	for idx, file := range files {
		path := fmt.Sprintf("%s/%s/%s/%s/%s/%s", repo, t.Channel, t.OS, repoVersion, arch, file)

		if repoVersion == "" {
			path = fmt.Sprintf("%s/%s/%s/%s/%s", repo, t.Channel, t.OS, arch, file)
		}

		files[idx] = path
	}

	return primaryPath, files, nil
}

func ParseRepoData(content string, attemps int) (string, []string, error) {
	type RepoMD struct {
		Data []struct {
			XMLName  xml.Name `xml:"data"`
			Type     string   `xml:"type,attr"`
			Location struct {
				XMLName xml.Name `xml:"location"`
				Href    string   `xml:"href,attr"`
			} `xml:"location"`
		} `xml:"data"`
	}

	var repoMD RepoMD
	if err := xml.Unmarshal([]byte(content), &repoMD); err != nil {
		attemps += 1

		if attemps > 2 {
			return "", nil, fmt.Errorf("xml.Unmarshal - %w", err)
		}

		time.Sleep(1 * time.Second)

		return ParseRepoData(content, attemps)
	}

	var primary string
	files := make([]string, 0)

	for _, md := range repoMD.Data {
		if md.Type == "primary" {
			primary = md.Location.Href
		}

		files = append(files, md.Location.Href)
	}

	return primary, files, nil
}

func RepoRelease(dist string) []string {
	files := []string{"Release", "InRelease"}
	var url []string

	for _, file := range files {
		url = append(url, fmt.Sprintf("%s/%s", dist, file))
	}

	return url
}
