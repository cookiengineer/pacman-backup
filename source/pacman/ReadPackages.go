package pacman

import "pacman-backup/structs"
import "os"

func ReadPackages(folder string) []structs.Package {

	var result []structs.Package

	stat, err1 := os.Stat(folder)

	if err1 == nil && stat.IsDir() {

		entries, err2 := os.ReadDir(folder)

		if err2 == nil {

			for e := 0; e < len(entries); e++ {

				file := entries[e].Name()

				if IsPackageFilename(file) {

					pkg := ReadPackage(folder + "/" + file)

					if pkg.Name != "" {
						result = append(result, pkg)
					}

				}

			}

		}

	}

	return result

}
