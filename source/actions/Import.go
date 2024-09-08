package actions

import "pacman-backup/console"
import "pacman-backup/pacman"
import "os"

func Import(sync_folder string, pkgs_folder string) bool {

	console.Group("Import")

	config := pacman.InitConfig()

	stat1, err1 := os.Stat(sync_folder)
	result_sync := true

	if err1 == nil && stat1.IsDir() {

		entries, err12 := os.ReadDir(sync_folder)

		if err12 == nil {

			for e := 0; e < len(entries); e++ {

				file := entries[e].Name()

				if pacman.IsDatabaseFilename(file) {

					buffer, err13 := os.ReadFile(sync_folder + "/" + file)

					if err13 == nil {

						err14 := os.WriteFile(config.Options.DBPath + "/sync/" + file, buffer, 0666)

						if err14 != nil {
							console.Error("> File \"sync/" + file + "\" not copied")
							result_sync = false
						}

					} else {
						console.Error("> File \"sync/" + file + "\" not copied")
						result_sync = false
					}

				}

			}

		}

	}

	stat2, err2 := os.Stat(pkgs_folder)
	result_pkgs := true

	if err2 == nil && stat2.IsDir() {

		entries, err22 := os.ReadDir(pkgs_folder)

		if err22 == nil {

			for e := 0; e < len(entries); e++ {

				file := entries[e].Name()

				if pacman.IsPackageFilename(file) {

					buffer, err23 := os.ReadFile(pkgs_folder + "/" + file)

					if err23 == nil {

						err24 := os.WriteFile(config.Options.CacheDir + "/" + file, buffer, 0666)

						if err24 != nil {
							console.Error("> File \"pkgs/" + file + "\" not copied")
							result_pkgs = false
						}

					} else {
						console.Error("> File \"pkgs/" + file + "\" not copied")
						result_pkgs = false
					}

				}

			}

		}

	}

	console.GroupEnd("Import")

	return result_sync && result_pkgs

}
