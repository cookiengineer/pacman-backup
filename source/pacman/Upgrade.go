package pacman

import "pacman-backup/sudo"
import "os"

func Upgrade(config string) (error) {

	var err error = nil

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	// Always needs sudo rights
	cmd1 := sudo.Command("pacman", "-Su", "--noconfirm", "--config", config)
	err1 := cmd1.Run()

	if err1 == nil {
		err = nil
	} else {
		err = err1
	}

	return err

}
