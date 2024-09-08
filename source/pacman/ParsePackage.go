package pacman

import "pacman-backup/matchers"
import "pacman-backup/structs"
import "pacman-backup/types"
import "strings"

func ParsePackage(buffer string, result *structs.Package) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var key string
	var val string

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.Contains(line, " : ") {

			key = strings.TrimSpace(line[0:strings.Index(line, " : ")])
			val = strings.TrimSpace(line[strings.Index(line, " : ")+3:])

		} else {

			val = strings.TrimSpace(line)

		}

		if key == "Architecture" {

			result.SetArchitecture(val)

		} else if key == "Build Date" {

			result.SetDatetime(val)

		} else if key == "Conflicts With" {

			if val != "None" {

				chunks := strings.Split(val, "  ")

				for c := 0; c < len(chunks); c++ {

					chunk := strings.TrimSpace(chunks[c])

					if chunk != "" {

						conflict := matchers.ToPackage(chunk)
						conflict.SetManager("pacman")
						result.AddConflict(conflict)

					}

				}

			}

		} else if key == "Depends On" {

			if val != "None" {

				chunks := strings.Split(val, "  ")

				for c := 0; c < len(chunks); c++ {

					chunk := strings.TrimSpace(chunks[c])

					if chunk != "" {

						dependency := matchers.ToPackage(chunk)
						dependency.SetManager("pacman")
						result.AddDependency(dependency)

					}

				}

			}

		} else if key == "Name" {

			result.SetName(val)

		} else if key == "Packager" {

			result.AddMaintainer(types.ToMaintainer(val))

		} else if key == "Provides" {

			if val != "None" {

				chunks := strings.Split(val, "  ")

				for c := 0; c < len(chunks); c++ {

					chunk := strings.TrimSpace(chunks[c])

					if chunk != "" {

						provide := matchers.ToPackage(chunk)
						provide.SetManager("pacman")
						result.AddProvide(provide)

					}

				}

			}

		} else if key == "Replaces" {

			if val != "None" {

				chunks := strings.Split(val, "  ")

				for c := 0; c < len(chunks); c++ {

					chunk := strings.TrimSpace(chunks[c])

					if chunk != "" {

						replace := matchers.ToPackage(chunk)
						replace.SetManager("pacman")
						result.AddReplace(replace)

					}

				}

			}

		} else if key == "URL" {

			if val != "None" {
				result.SetURL(val)
			}

		} else if key == "Version" {

			if val != "None" {
				result.SetVersion(val)
			}

		}

	}

}
