package matchers

import "pacman-backup/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Package struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Vendor       string `json:"vendor"`
}

func NewPackage() Package {

	var pkg Package

	pkg.Version = "any"
	pkg.Architecture = "any"
	pkg.Manager = "any"
	pkg.Vendor = "any"

	return pkg

}

func ToPackage(value string) Package {

	var pkg Package

	pkg.Version = "any"
	pkg.Architecture = "any"
	pkg.Manager = "any"
	pkg.Vendor = "any"

	pkg.Parse(value)

	return pkg

}

func (pkg *Package) IsIdentical(value Package) bool {

	var result bool = false

	if pkg.Name == value.Name &&
		pkg.Version == value.Version &&
		pkg.Architecture == value.Architecture &&
		pkg.Manager == value.Manager {
		result = true
	}

	return result

}

func (pkg *Package) IsValid() bool {

	var result bool = false

	if pkg.Name != "" {
		result = true
	}

	return result

}

func (pkg *Package) Matches(name string, version string, manager string, vendor string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(version, " ") {
		version = strings.TrimSpace(version[strings.Index(version, " ")+1:])
	}

	var matches_name bool = false
	var matches_version bool = false
	var matches_manager bool = false
	var matches_vendor bool = false

	if pkg.Name == name {
		matches_name = true
	} else if pkg.Name == "any" {
		matches_name = true
	}

	if pkg.Version == "any" {

		matches_version = true

	} else if strings.HasPrefix(pkg.Version, "<= ") {

		pkg_version := types.ToVersion(pkg.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(pkg_version) {
			matches_version = true
		} else if other_version.IsBefore(pkg_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(pkg.Version, "< ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsBefore(pkg_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(pkg.Version, ">= ") {

		pkg_version := types.ToVersion(pkg.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(pkg_version) {
			matches_version = true
		} else if other_version.IsAfter(pkg_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(pkg.Version, "> ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsAfter(pkg_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(pkg.Version, "= ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(pkg_version) {
			matches_version = true
		}

	} else {

		pkg_version := types.ToVersion(pkg.Version)
		other_version := types.ToVersion(version)

		if other_version.IsSame(pkg_version) {
			matches_version = true
		}

	}

	if pkg.Manager == manager {
		matches_manager = true
	} else if pkg.Manager == "any" {
		matches_manager = true
	}

	if pkg.Vendor == vendor {
		matches_vendor = true
	} else if pkg.Vendor == "any" {
		matches_vendor = true
	}

	return matches_name && matches_version && matches_manager && matches_vendor

}

func (pkg *Package) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	pkg.Name = name
	pkg.Version = version

	if architecture != "" {
		pkg.Architecture = architecture
	}

}

func (pkg *Package) SetArchitecture(value string) {
	pkg.Architecture = value
}

func (pkg *Package) SetManager(value string) {
	pkg.Manager = value
}

func (pkg *Package) SetName(value string) {
	pkg.Name = strings.TrimSpace(value)
}

func (pkg *Package) SetVendor(value string) {

	if value == "all" || value == "any" || value == "*" {
		pkg.Vendor = "any"
	} else if value != "" {
		pkg.Vendor = value
	}

}

func (pkg *Package) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		pkg.Version = "any"
	} else if value != "" {
		pkg.Version = value
	}

}

func (pkg *Package) Hash() string {

	var hash string

	if pkg.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			pkg.Name,
			pkg.Version,
			pkg.Architecture,
			pkg.Manager,
			pkg.Vendor,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
