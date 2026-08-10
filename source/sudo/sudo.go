package sudo

import "os"
import "sync"

var sudo_password string
var sudo_mutex    sync.Mutex

func SetPassword(password string) {

	sudo_mutex.Lock()
	defer sudo_mutex.Unlock()

	sudo_password = password

}

func GetPassword() string {

	sudo_mutex.Lock()
	defer sudo_mutex.Unlock()

	return sudo_password

}

func IsRoot() bool {
	return os.Geteuid() == 0
}

func NeedsSudo() bool {

	if IsRoot() == true {

		return false

	} else {

		if GetPassword() == "" {
			return true
		} else {
			return false
		}

	}

}

