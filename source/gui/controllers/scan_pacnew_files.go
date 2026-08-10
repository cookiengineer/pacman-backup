package controllers

import "io/fs"
import "path/filepath"
import "strings"

func scan_pacnew_files() []string {

	files := make([]string, 0)

	filepath.WalkDir("/etc", func(path string, dir_entry fs.DirEntry, err error) error {

		if err != nil {
			return nil
		}

		if dir_entry.IsDir() {
			return nil
		}

		if strings.HasSuffix(dir_entry.Name(), ".pacnew") || strings.HasSuffix(dir_entry.Name(), ".pacsave") {
			files = append(files, path)
		}

		return nil

	})

	return files

}
