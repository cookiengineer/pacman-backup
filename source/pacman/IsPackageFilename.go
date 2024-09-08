package pacman

import "strings"

func IsPackageFilename(filepath string) bool {

	var result bool = false

	filename := ""

	if strings.Contains(filepath, "/") {
		filename = filepath[strings.LastIndex(filepath, "/")+1:]
	} else {
		filename = filepath
	}

	if strings.HasSuffix(filename, ".pkg.tar.xz") {
		result = true
	} else if strings.HasSuffix(filename, ".pkg.tar.zst") {
		result = true
	}

	return result

}

