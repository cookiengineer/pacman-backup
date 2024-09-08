package pacman

import "pacman-backup/structs"
import "os"

func CollectFiles(config string, folder string) []structs.Package {

	var result []structs.Package

	stat, err1 := os.Stat(folder)

	if err1 == nil && stat.IsDir() {

		entries, err2 := os.ReadDir(folder)

		if err2 == nil {

			for e := 0; e < len(entries); e++ {

				file := entries[e].Name()

				if IsPackageFilename(file) {

					pkg := CollectFile(config, folder + "/" + file)

					if pkg.Name != "" {
						result = append(result, pkg)
					}

				}

			}

		}

	}

	return result

}
