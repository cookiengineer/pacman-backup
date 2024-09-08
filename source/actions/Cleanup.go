package actions

import "pacman-backup/pacman"
import "pacman-backup/structs"

func Cleanup(pkgs_folder string) bool {

	var result bool

	database := structs.NewDatabase()
	packages := pacman.ReadPackages(pkgs_folder)

	if len(packages) > 0 {

		for p := 0; p < len(packages); p++ {

			pkg := packages[p]

			if pkg.IsValid() {
				database.AddPackage(pkg)
			}

			// TODO: map packages to their architecture (any, armv7h, aarch64, i686, x86_64)

		}

	}

	return result

}
