package actions

import "pacman-backup/console"
import "pacman-backup/pacman"
import "os"

func Export(sync_folder string, pkgs_folder string) bool {

	console.Group("Export")

	config := pacman.InitConfig()

	stat1, err1 := os.Stat(config.Options.DBPath + "/sync")
	result_sync := true

	if err1 == nil && stat1.IsDir() {

		entries, err12 := os.ReadDir(config.Options.DBPath + "/sync")

		if err12 == nil {

			for e := 0; e < len(entries); e++ {

				file := entries[e].Name()

				if pacman.IsDatabaseFilename(file) {

					buffer, err13 := os.ReadFile(config.Options.DBPath + "/sync/" + file)

					if err13 == nil {

						err14 := os.WriteFile(sync_folder + "/" + file, buffer, 0666)

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

	stat2, err2 := os.Stat(config.Options.CacheDir)
	result_pkgs := true

	if err2 == nil && stat2.IsDir() {

		entries, err22 := os.ReadDir(config.Options.CacheDir)

		if err22 == nil {

			for e := 0; e < len(entries); e++ {

				file := entries[e].Name()

				if pacman.IsPackageFilename(file) {

					buffer, err23 := os.ReadFile(config.Options.CacheDir + "/" + file)

					if err23 == nil {

						err24 := os.WriteFile(pkgs_folder + "/" + file, buffer, 0666)

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

	console.GroupEnd("Export")

	return result_sync && result_pkgs

}
