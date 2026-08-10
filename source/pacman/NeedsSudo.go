package pacman

import "os"
import "golang.org/x/sys/unix"

func can_read_file(file string) bool {

	info, err := os.Stat(file)

	if err == nil && info.IsDir() == false {
		return unix.Access(file, unix.R_OK) == nil
	} else {
		return false
	}

}

func can_write_file(file string) bool {

	info, err := os.Stat(file)

	if err == nil && info.IsDir() == false {
		return unix.Access(file, unix.W_OK) == nil
	} else {
		return false
	}

}

func can_write_folder(folder string) bool {

	info, err := os.Stat(folder)

	if err == nil && info.IsDir() == true {
		return unix.Access(folder, unix.W_OK|unix.X_OK) == nil
	} else {
		return false
	}

}


func NeedsSudo(config_path string) bool {

	if can_read_file(config_path) == true {

		config := InitConfig(config_path)

		if can_write_folder(config.Options.DBPath) && can_write_folder(config.Options.CacheDir) && can_write_file(config.Options.LogFile) {
			return false
		} else {
			return true
		}

	} else {
		return false
	}

}
