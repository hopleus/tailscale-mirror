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

		if osTrack.PackageType == "yum" || osTrack.PackageType == "dnf" {
			for _, arch := range ArchList {
				repoMd, dist := RepoMD(osTrack, arch)
				DownloadFiles([]string{repoMd, fmt.Sprintf("%s.asc", repoMd)}, true)

				primary, files, err := RepoData(osTrack, arch, repoMd)
				if err != nil {
					fmt.Println("[Error] RepoData -", err)
					continue
				}

				DownloadFiles(files, true)
				rpm, err := PoolRpmFromPrimaryMD(dist, primary)
				if err != nil {
					fmt.Println("[Error] PoolRpmFromPrimaryMD: %w", err)
					continue
				}

				debs = append(debs, rpm...)
			}
		} else {
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
		}

		GenerateDoc(osTrack)
	}

	debs = UniqueArray(debs)
	DownloadFiles(debs, true)
}
