package actions

import "pacman-backup/pacman"
import "pacman-backup/structs"
import "os"

func Export(console *structs.Console, sync_folder string, pkgs_folder string) bool {

	console.Group("actions/Export")

	config := pacman.InitConfig()

	stat1, err1 := os.Stat(config.Options.DBPath + "/sync")
	result_sync := true

	if err1 == nil && stat1.IsDir() {

		entries, err12 := os.ReadDir(config.Options.DBPath + "/sync")

		if err12 == nil {

			for e := 0; e < len(entries); e++ {

				filename := entries[e].Name()

				if pacman.IsDatabaseFilename(filename) {

					console.Progress("File sync/" + filename)

					buffer, err13 := os.ReadFile(config.Options.DBPath + "/sync/" + filename)

					if err13 == nil {

						err14 := os.WriteFile(sync_folder + "/" + filename, buffer, 0666)

						if err14 != nil {
							console.Error("File sync/" + filename + " failed to copy")
							result_sync = false
						}

					} else {
						console.Error("File sync/" + filename + " failed to copy")
						result_sync = false
					}

				}

			}

		}

	}

	stat2, err2 := os.Stat(config.Options.CacheDir)
	result_pkgs := true

	if err2 == nil && stat2.IsDir() {

		entries, err22 := os.ReadDir(config.Options.CacheDir)

		if err22 == nil {

			for e := 0; e < len(entries); e++ {

				filename := entries[e].Name()

				if pacman.IsPackageFilename(filename) {

					console.Progress("File pkgs/" + filename)

					buffer, err23 := os.ReadFile(config.Options.CacheDir + "/" + filename)

					if err23 == nil {

						err24 := os.WriteFile(pkgs_folder + "/" + filename, buffer, 0666)

						if err24 != nil {
							console.Error("File pkgs/" + filename + " failed to copy")
							result_pkgs = false
						}

					} else {
						console.Error("File pkgs/" + filename + " failed to copy")
						result_pkgs = false
					}

				}

			}

		}

	}

	if result_sync && result_pkgs {
		console.GroupEnd("actions/Export succeeded")
	} else {
		console.GroupEnd("actions/Export failed")
	}

	return result_sync && result_pkgs

}
