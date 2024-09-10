package pacman

import "strings"

func IsDatabaseFilename(filepath string) bool {

	var result bool = false

	filename := ""

	if strings.Contains(filepath, "/") {
		filename = filepath[strings.LastIndex(filepath, "/")+1:]
	} else {
		filename = filepath
	}

	if strings.HasSuffix(filename, ".db") {
		result = true
	} else if strings.HasSuffix(filename, ".db.sig") {
		result = true
	} else if strings.HasSuffix(filename, ".files") {
		result = true
	}

	return result

}

