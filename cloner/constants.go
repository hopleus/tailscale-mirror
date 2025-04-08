package main

var OS = []string{"ubuntu", "debian", "raspian", "centos", "fedora", "rhel", "amazon-linux", "oracle", "photon"}
var OSTracks = []OSTrack{
	// ubuntu (legacy)
	{
		OS:          "ubuntu",
		Version:     "xenial",
		PackageType: "apt",
		AptKeyType:  "legacy",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "bionic",
		PackageType: "apt",
		AptKeyType:  "legacy",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "eoan",
		PackageType: "apt",
		AptKeyType:  "legacy",
		Channel:     "stable",
	},
	// ubuntu (keyring)
	{
		OS:          "ubuntu",
		Version:     "focal",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "groovy",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "hirsute",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "impish",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "jammy",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "juniper",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "kinetic",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "lunar",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "mantic",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "minotaur",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "ubuntu",
		Version:     "noble",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "debian",
		Version:     "stretch",
		PackageType: "apt",
		AptKeyType:  "legacy",
		Channel:     "stable",
	},
	{
		OS:          "debian",
		Version:     "buster",
		PackageType: "apt",
		AptKeyType:  "legacy",
		Channel:     "stable",
	},
	{
		OS:          "debian",
		Version:     "bullseye",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "debian",
		Version:     "bookworm",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "debian",
		Version:     "sid",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "debian",
		Version:     "trixie",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "raspbian",
		Version:     "stretch",
		PackageType: "apt",
		AptKeyType:  "legacy",
		Channel:     "stable",
	},
	{
		OS:          "raspbian",
		Version:     "buster",
		PackageType: "apt",
		AptKeyType:  "legacy",
		Channel:     "stable",
	},
	{
		OS:          "raspbian",
		Version:     "bullseye",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	{
		OS:          "raspbian",
		Version:     "bookworm",
		PackageType: "apt",
		AptKeyType:  "keyring",
		Channel:     "stable",
	},
	// TODO: Add centos, fedora, etc... in the future
	// centos (yum)
	// {
	// 	OS:          "centos",
	// 	Version:     "7",
	// 	PackageType: "yum",
	// 	Channel:     "stable",
	// },
	// centos (dnf)
	// {
	// 	OS:          "centos",
	// 	Version:     "8",
	// 	PackageType: "dnf",
	// 	Channel:     "stable",
	// },
}

const (
	repo       = "https://pkgs.tailscale.com"
	mirror     = "https://raw.githubusercontent.com/hopleus/tailscale-mirror/main/data"
	dataDir    = "../data"
	docDir     = "../docs"
	stubDir    = "../stubs"
	minVersion = "1.82.0"

	regExpReleasePackagePattern = `\w{32}\s\d+(.*Packages(\.gz)?)\s`
	regExpPackageSectionPattern = `(?m)(\n\n|\n$)`
	regExpPackageVersionPattern = `(?m)Version:\s(.*)`
	regExpPackagePoolPattern    = `(?m)Filename:\s(.*)`
)
