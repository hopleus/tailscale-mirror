package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ReleasePackages(dist, release string) ([]string, error) {
	path := LocalRepoData(release)
	content, err := GetContentFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("GetContentFromFile - %w", err)
	}

	pkgs := ParseReleaseFile(content)
	var distPkgs []string

	for _, pkg := range pkgs {
		distPkgs = append(
			distPkgs,
			fmt.Sprintf(
				"%s/%s",
				dist,
				strings.TrimSpace(pkg),
			),
		)
	}

	DownloadFiles(distPkgs, true)

	return pkgs, nil
}

func ParseReleaseFile(content string) []string {
	pattern := regexp.MustCompile(regExpReleasePackagePattern)
	matches := pattern.FindAllStringSubmatch(content, -1)

	packages := make(map[string]bool)
	uniquePackages := make([]string, 0)
	var packageUrl string

	for _, match := range matches {
		packageUrl = match[1]
		if ok := packages[packageUrl]; !ok {
			packages[packageUrl] = true
			uniquePackages = append(uniquePackages, packageUrl)
		}
	}

	return uniquePackages
}

func SectionsPackageFile(content string) []string {
	pattern := regexp.MustCompile(regExpPackageSectionPattern)

	return pattern.Split(content, -1)
}

func PoolFromPackages(distUrl, repoUrl string, packages []string) []string {
	var debs []string

	var distPackages []string
	for _, packageName := range packages {
		distPackages = append(
			distPackages,
			fmt.Sprintf(
				"%s/%s",
				distUrl,
				strings.TrimSpace(packageName),
			),
		)
	}

	var localPackagePath string

	for _, dist := range distPackages {
		localPackagePath = LocalRepoData(strings.TrimSpace(dist))
		content, err := GetContentFromFile(localPackagePath)
		if err != nil {
			continue
		}

		var allowed bool
		var allowedSections []string
		packageSections := SectionsPackageFile(content)
		for _, section := range packageSections {
			if strings.Contains(section, "Package: tailscale\n") {
				allowed = CheckPackageVersion(section)
				if !allowed {
					continue
				}
			}

			allowedSections = append(allowedSections, section)
		}

		sections := fmt.Sprintf("%s\n", strings.Join(allowedSections, "\n\n"))
		sections = strings.Replace(sections, "\n\n\n", "\n\n", -1)

		debsUrl := DebsPoolPackages(sections)

		for _, deb := range debsUrl {
			debs = append(
				debs,
				fmt.Sprintf(
					"%s/%s",
					repoUrl,
					strings.TrimSpace(deb),
				),
			)
		}
	}

	return debs
}

func CheckPackageVersion(content string) bool {
	pattern := regexp.MustCompile(regExpPackageVersionPattern)
	versionFound := pattern.FindAllStringSubmatch(content, -1)
	if len(versionFound) == 0 {
		return false
	}

	version := versionFound[0][1]
	simplifyMinVersion := VersionBytes(minVersion)
	simplifyVersion := VersionBytes(version)

	return simplifyVersion >= simplifyMinVersion
}

func DebsPoolPackages(pkg string) []string {
	var debs []string

	pattern := regexp.MustCompile(regExpPackagePoolPattern)
	debsFound := pattern.FindAllStringSubmatch(pkg, -1)
	if len(debsFound) == 0 {
		return debs
	}

	for _, deb := range debsFound {
		if len(deb) == 0 {
			continue
		}

		debs = append(debs, deb[1])
	}

	return debs
}
