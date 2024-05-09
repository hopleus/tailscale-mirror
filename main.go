package main

import "fmt"

func main() {
	var debs []string

	for _, osTrack := range OSTracks {
		url := osTrack.UrlRepo()
		DownloadFiles(url, true)

		if err := ReplaceSourceRepoToMirror(url[0]); err != nil {
			fmt.Println("\n[Error] ReplaceSourceRepoToMirror %w", err)
			continue
		}

		repo, dist := RepoDist(osTrack)
		release := RepoRelease(dist)
		DownloadFiles(release, true)

		for _, _release := range release {
			packages, err := ReleasePackages(dist, _release)
			if err != nil {
				fmt.Println("[Error] ReleasePackages: %w", err)
				continue
			}

			debs = append(debs, PoolFromPackages(dist, repo, packages)...)
		}

		GenerateDoc(osTrack)
	}

	debs = UniqueArray(debs)
	DownloadFiles(debs, true)
}
