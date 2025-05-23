package main

import "fmt"

type OSTrack struct {
	OS          string
	Version     string
	PackageType string
	AptKeyType  string
	Channel     string
}

func (t *OSTrack) Stub() string {
	if t.PackageType == "apt" {
		return fmt.Sprintf("%s/apt.md", stubDir)
	}

	if t.PackageType == "yum" {
		return fmt.Sprintf("%s/yum.md", stubDir)
	}

	if t.PackageType == "dnf" {
		return fmt.Sprintf("%s/dnf.md", stubDir)
	}

	return ""
}

func (t *OSTrack) UrlRepo() []string {
	path := fmt.Sprintf("%s/%s/%s/%s", repo, t.Channel, t.OS, t.Version)

	if t.Version == "" {
		path = fmt.Sprintf("%s/%s/%s", repo, t.Channel, t.OS)
	}

	if InArray([]string{"ubuntu", "debian", "raspbian"}, t.OS) {
		if t.AptKeyType == "legacy" {
			return []string{
				fmt.Sprintf("%s.list", path),
				fmt.Sprintf("%s.asc", path),
			}
		}

		return []string{
			fmt.Sprintf("%s.tailscale-keyring.list", path),
			fmt.Sprintf("%s.noarmor.gpg", path),
		}
	}

	// TODO: Add in the future
	if InArray([]string{"centos", "fedora", "rhel", "amazon-linux"}, t.OS) {
		return []string{
			fmt.Sprintf("%s/tailscale.repo", path),
			fmt.Sprintf("%s/repo.gpg", path),
		}
	}

	return []string{}
}
