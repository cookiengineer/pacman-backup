package structs

import "pacman-backup/matchers"
import "pacman-backup/types"
import "strings"

type Package struct {
	Name         string                `json:"name"`
	Version      types.Version         `json:"version"`
	Architecture string                `json:"architecture"`
	Manager      string                `json:"manager"`
	Vendor       string                `json:"vendor"`
	URL          string                `json:"url"`
	Datetime     types.Datetime        `json:"datetime"`
	Maintainers  []types.Maintainer    `json:"maintainers"`
	Filesystem   []string              `json:"filesystem"`
	Conflicts    []matchers.Package    `json:"conflicts"`
	Dependencies []matchers.Package    `json:"dependencies"`
	Provides     []matchers.Package    `json:"provides"`
	Replaces     []matchers.Package    `json:"replaces"`
	Unresolved   []matchers.Unresolved `json:"unresolved"`
}

func NewPackage(manager string) Package {

	var pkg Package

	pkg.SetManager(manager)

	pkg.Maintainers = make([]types.Maintainer, 0)
	pkg.Filesystem = make([]string, 0)

	pkg.Conflicts = make([]matchers.Package, 0)
	pkg.Dependencies = make([]matchers.Package, 0)
	pkg.Provides = make([]matchers.Package, 0)
	pkg.Replaces = make([]matchers.Package, 0)
	pkg.Unresolved = make([]matchers.Unresolved, 0)

	return pkg

}

func (pkg *Package) IsIdentical(value Package) bool {

	var result bool = false

	if pkg.Name == value.Name &&
		pkg.Version.String() == value.Version.String() &&
		pkg.Architecture == value.Architecture &&
		pkg.Manager == value.Manager {
		result = true
	}

	return result

}

func (pkg *Package) IsValid() bool {

	if pkg.Name != "" {

		if pkg.Architecture != "" && pkg.Manager != "" {

			var result bool = true

			if pkg.Datetime.IsValid() == false {
				result = false
			}

			if pkg.Version.IsValid() == false {
				result = false
			}

			if result == true {

				for m := 0; m < len(pkg.Maintainers); m++ {

					if pkg.Maintainers[m].IsValid() == false {
						result = false
						break
					}

				}

			}

			if result == true {

				for c := 0; c < len(pkg.Conflicts); c++ {

					if pkg.Conflicts[c].IsValid() == false {
						result = false
						break
					}

				}

			}

			if result == true {

				for d := 0; d < len(pkg.Dependencies); d++ {

					if pkg.Dependencies[d].IsValid() == false {
						result = false
						break
					}

				}

			}

			if result == true {

				for p := 0; p < len(pkg.Provides); p++ {

					if pkg.Provides[p].IsValid() == false {
						result = false
						break
					}

				}

			}

			if result == true {

				for r := 0; r < len(pkg.Replaces); r++ {

					if pkg.Replaces[r].IsValid() == false {
						result = false
						break
					}

				}

			}

			return result

		}

	}

	return false

}

func (pkg *Package) SetArchitecture(value string) {
	pkg.Architecture = value
}

func (pkg *Package) AddConflict(value matchers.Package) {

	if value.IsValid() {

		var found bool = false

		for c := 0; c < len(pkg.Conflicts); c++ {

			if pkg.Conflicts[c].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Conflicts = append(pkg.Conflicts, value)
		}

	}

}

func (pkg *Package) RemoveConflict(value matchers.Package) {

	var index int = -1

	for c := 0; c < len(pkg.Conflicts); c++ {

		if pkg.Conflicts[c].IsIdentical(value) {
			index = c
			break
		}

	}

	if index != -1 {
		pkg.Conflicts = append(pkg.Conflicts[:index], pkg.Conflicts[index+1:]...)
	}

}

func (pkg *Package) SetConflicts(value []matchers.Package) {

	var filtered []matchers.Package

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Conflicts = filtered

}

func (pkg *Package) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		pkg.Datetime = datetime
	}

}

func (pkg *Package) AddDependency(value matchers.Package) {

	if value.IsValid() {

		var found bool = false

		for d := 0; d < len(pkg.Dependencies); d++ {

			if pkg.Dependencies[d].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Dependencies = append(pkg.Dependencies, value)
		}

	}

}

func (pkg *Package) HasDependency(dependency matchers.Package) bool {

	var result bool = false

	for d := 0; d < len(pkg.Dependencies); d++ {

		other := pkg.Dependencies[d]

		if dependency.Name == other.Name &&
			dependency.Version == other.Version {
			result = true
			break
		}

	}

	return result

}

func (pkg *Package) RemoveDependency(value matchers.Package) {

	var index int = -1

	for d := 0; d < len(pkg.Dependencies); d++ {

		if pkg.Dependencies[d].IsIdentical(value) {
			index = d
			break
		}

	}

	if index != -1 {
		pkg.Dependencies = append(pkg.Dependencies[:index], pkg.Dependencies[index+1:]...)
	}

}

func (pkg *Package) ResolveDependencies(packages []Package) {

	if len(pkg.Unresolved) > 0 && len(packages) > 0 {

		var remaining []matchers.Unresolved

		for u := 0; u < len(pkg.Unresolved); u++ {

			unresolved := pkg.Unresolved[u]

			var resolved matchers.Package

			for p := 0; p < len(packages); p++ {

				other := packages[p]

				if unresolved.Matches(other.Name, other.Version.String(), other.Manager, other.Vendor) {

					resolved.Name = other.Name
					resolved.Version = other.Version.String()

					if other.Manager != "" {
						resolved.Manager = other.Manager
					} else {
						resolved.Manager = "any"
					}

					if other.Vendor != "" {
						resolved.Vendor = other.Vendor
					} else {
						resolved.Vendor = "any"
					}

				} else {

					for p := 0; p < len(other.Provides); p++ {

						provide := other.Provides[p]

						if unresolved.Matches(provide.Name, provide.Version, provide.Manager, provide.Vendor) {

							resolved.Name = other.Name
							resolved.Version = other.Version.String()

							if other.Manager != "" {
								resolved.Manager = other.Manager
							} else {
								resolved.Manager = "any"
							}

							if other.Vendor != "" {
								resolved.Vendor = other.Vendor
							} else {
								resolved.Vendor = "any"
							}

						}

					}

				}

				if resolved.Name != "" {
					break
				}

			}

			if resolved.Name != "" {

				if !pkg.HasDependency(resolved) {
					pkg.AddDependency(resolved)
				}

			} else {

				remaining = append(remaining, unresolved)

			}

		}

		pkg.Unresolved = remaining

	}

}

func (pkg *Package) SetDependencies(value []matchers.Package) {

	var filtered []matchers.Package

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Dependencies = filtered

}

func (pkg *Package) HasFile(value string) bool {

	var result bool = false

	for f := 0; f < len(pkg.Filesystem); f++ {

		if pkg.Filesystem[f] == value {
			result = true
			break
		}

	}

	return result

}

func (pkg *Package) SetFilesystem(value []string) {

	var filtered []string

	for v := 0; v < len(value); v++ {

		file := value[v]
		found := false

		for f := 0; f < len(filtered); f++ {

			if filtered[f] == file {
				found = true
				break
			}

		}

		if found == false {
			filtered = append(filtered, file)
		}

	}

	pkg.Filesystem = filtered

}

func (pkg *Package) AddMaintainer(value types.Maintainer) {

	if value.IsValid() {

		var found bool = false

		for m := 0; m < len(pkg.Maintainers); m++ {

			if pkg.Maintainers[m].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Maintainers = append(pkg.Maintainers, value)
		}

	}

}

func (pkg *Package) RemoveMaintainer(value types.Maintainer) {

	var index int = -1

	for m := 0; m < len(pkg.Maintainers); m++ {

		if pkg.Maintainers[m].IsIdentical(value) {
			index = m
			break
		}

	}

	if index != -1 {
		pkg.Maintainers = append(pkg.Maintainers[:index], pkg.Maintainers[index+1:]...)
	}

}

func (pkg *Package) SetMaintainers(value []types.Maintainer) {

	var filtered []types.Maintainer

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Maintainers = filtered

}

func (pkg *Package) SetManager(value string) {
	pkg.Manager = value
}

func (pkg *Package) SetName(value string) {
	pkg.Name = strings.TrimSpace(value)
}

func (pkg *Package) AddProvide(value matchers.Package) {

	if value.IsValid() {

		var found bool = false

		for p := 0; p < len(pkg.Provides); p++ {

			if pkg.Provides[p].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Provides = append(pkg.Provides, value)
		}

	}

}

func (pkg *Package) RemoveProvide(value matchers.Package) {

	var index int = -1

	for p := 0; p < len(pkg.Provides); p++ {

		if pkg.Provides[p].IsIdentical(value) {
			index = p
			break
		}

	}

	if index != -1 {
		pkg.Provides = append(pkg.Provides[:index], pkg.Provides[index+1:]...)
	}

}

func (pkg *Package) SetProvides(value []matchers.Package) {

	var filtered []matchers.Package

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Provides = filtered

}

func (pkg *Package) AddReplace(value matchers.Package) {

	if value.IsValid() {

		var found bool = false

		for r := 0; r < len(pkg.Replaces); r++ {

			if pkg.Replaces[r].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			pkg.Replaces = append(pkg.Replaces, value)
		}

	}

}

func (pkg *Package) RemoveReplace(value matchers.Package) {

	var index int = -1

	for r := 0; r < len(pkg.Replaces); r++ {

		if pkg.Replaces[r].IsIdentical(value) {
			index = r
			break
		}

	}

	if index != -1 {
		pkg.Replaces = append(pkg.Replaces[:index], pkg.Replaces[index+1:]...)
	}

}

func (pkg *Package) SetReplaces(value []matchers.Package) {

	var filtered []matchers.Package

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	pkg.Replaces = filtered

}

func (pkg *Package) SetURL(value string) {
	pkg.URL = value
}

func (pkg *Package) SetVendor(value string) {
	pkg.Vendor = strings.TrimSpace(value)
}

func (pkg *Package) SetVersion(value string) {

	version := types.ToVersion(value)

	if version.IsValid() {
		pkg.Version = version
	}

}
