package main

import (
	"fmt"
	"os"
	"strings"
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

	return path, fmt.Sprintf("%s/dists/%s", path, t.Version)
}

func RepoRelease(dist string) []string {
	files := []string{"Release", "InRelease"}
	var url []string

	for _, file := range files {
		url = append(url, fmt.Sprintf("%s/%s", dist, file))
	}

	return url
}
