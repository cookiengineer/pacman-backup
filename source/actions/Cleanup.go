package actions

import "pacman-backup/console"
import "pacman-backup/pacman"
import "pacman-backup/types"
import "os"
import "strconv"

func Cleanup(sync_folder string, pkgs_folder string) bool {

	console.Group("actions/Cleanup")

	result := false
	config := pacman.NewConfig("", sync_folder, pkgs_folder)
	err1 := os.WriteFile("/tmp/pacman-backup.conf", []byte(config.String()), 0666)

	if err1 == nil {

		packages := pacman.CollectFiles("/tmp/pacman-backup.conf", pkgs_folder)
		result = true

		if len(packages) > 0 {

			console.Log("Found " + strconv.Itoa(len(packages)) + " Packages")

			package_to_filenames := make(map[string][]string, 0)
			package_to_versions := make(map[string][]types.Version, 0)

			for filename, pkg := range packages {


				_, ok1 := package_to_filenames[pkg.Name]

				if ok1 == true {
					package_to_filenames[pkg.Name] = append(package_to_filenames[pkg.Name], filename)
				} else {
					package_to_filenames[pkg.Name] = []string{filename}
				}

				_, ok2 := package_to_versions[pkg.Name]

				if ok2 == true {
					package_to_versions[pkg.Name] = append(package_to_versions[pkg.Name], pkg.Version)
				} else {
					package_to_versions[pkg.Name] = []types.Version{pkg.Version}
				}

			}

			for pkgname, versions := range package_to_versions {

				console.Progress("Package " + pkgname)

				if len(versions) > 1 {

					var last_version_index int = -1

					for v := 0; v < len(versions); v++ {

						if last_version_index == -1 {
							last_version_index = v
						} else {

							old_version := versions[last_version_index]
							new_version := versions[v]

							if new_version.IsAfter(old_version) {
								last_version_index = v
							}

						}

					}

					if last_version_index != -1 {

						console.Group("Package " + pkgname + " has " + strconv.Itoa(len(versions) - 1) + " old versions")

						for v := 0; v < len(versions); v++ {

							if v != last_version_index {

								filename := package_to_filenames[pkgname][v]

								err2 := os.Remove(pkgs_folder + "/" + filename)

								if err2 != nil {
									console.Warn("File " + filename + " failed to remove")
								}

							}

						}

						console.GroupEnd("Package " + pkgname)

					}

				}

			}

		} else {

			console.Warn("Found 0 Packages")
			result = true

		}

	} else {
		console.Error(err1.Error())
	}

	console.GroupEndResult(result, "actions/Cleanup")

	return result

}
