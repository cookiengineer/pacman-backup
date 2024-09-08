package main

import "pacman-backup/actions"
import "pacman-backup/console"
import "pacman-backup/pacman"
import "os"
import "strings"

func isFolder(value string) bool {

	if strings.HasPrefix(value, "/") {

		stat, err := os.Stat(value)

		if err == nil && stat.IsDir() {
			return true
		}

	}

	return false

}

func makeFolder(value string) bool {

	err := os.MkdirAll(value, 0666)

	if err == nil {
		return true
	}

	return false

}

func isMirror(value string) bool {

	var result bool

	if strings.HasPrefix(value, "http://") && strings.HasSuffix(value, ":15678") {
		result = true
	}

	return result

}

func showUsage() {

	user := os.Getenv("USER")

	if user == "root" {
		user = os.Getenv("SUDO_USER")
	}

	console.Info("")
	console.Info("Pacman Backup")
	console.Info("Offline Pacman Cache Management Tool")
	console.Info("")

	console.Group("Usage: tholian-guard [Action] [Folder]")
	console.Log("")
	console.Log("The [Folder] parameter is optional. If no folder is set, pacman's default folders will be used.")
	console.Log("(/var/lib/pacman/sync and /var/cache/pacman/pkg)")
	console.Log("")
	console.GroupEnd("------")

	console.Group("Action   | Description                                                                 |")
	console.Log("---------|-----------------------------------------------------------------------------|")
	console.Log("export   | Exports packages and package index to a specified folder.                   |")
	console.Log("import   | Imports packages and package index from a specified folder.                 |")
	console.Log("---------|-----------------------------------------------------------------------------|")
	console.Log("cleanup  | Cleans up outdated packages from a specified folder.                        |")
	console.Log("upgrade  | Upgrades all upgradable packages.                                           |")
	console.Log("---------|-----------------------------------------------------------------------------|")
	console.Log("download | Downloads all available packages and package index from a specified mirror. |")
	console.Log("serve    | Serves a local mirror from a specified folder.                              |")
	console.GroupEnd("---------|-----------------------------------------------------------------------------|")

	console.Group("USB Drive Example")
	console.Log("# Step 1: Machine with internet connection")
	console.Log("> sudo pacman-backup download;")
	console.Log("> pacman-backup export /run/media/" + user + "/pacman-usbdrive;")
	console.Log("> pacman-backup cleanup /run/media/" + user + "/pacman-usbdrive;")
	console.Log("> sync;")
	console.Log("")
	console.Log("# Step 2: Machine without internet connection")
	console.Log("> sudo pacman-backup import /run/media/" + user + "/pacman-usbdrive;")
	console.Log("> sudo pacman-backup upgrade /run/media/" + user + "/pacman-usbdrive;")
	console.Log("> sync;")
	console.GroupEnd("-----------------")

	console.Group("LAN Mirror Example")
	console.Log("# Step 1: Machine with internet connection")
	console.Log("> sudo pacman-backup download;")
	console.Log("> sudo pacman-backup serve;")
	console.Log("")
	console.Log("# Step 2: Machine without internet connection")
	console.Log("> sudo pacman-backup download http://192.168.0.10:15678")
	console.Log("> sudo pacman-backup upgrade;")
	console.Group("------------------")

}

func main() {

	if len(os.Args) == 4 {

		action := os.Args[1]

		if action == "download" {

			// pacman-backup download http://mirror:15678 /mnt/usb-drive
			if isMirror(os.Args[2]) && isFolder(os.Args[3]) {

				mirror := os.Args[2]

				if !isFolder(os.Args[3] + "/sync") {
					makeFolder(os.Args[3] + "/sync")
				}

				if !isFolder(os.Args[3] + "/pkgs") {
					makeFolder(os.Args[3] + "/pkgs")
				}

				actions.Sync(mirror, os.Args[3] + "/sync", os.Args[3] + "/pkgs")
				actions.Download(mirror, os.Args[3] + "/sync", os.Args[3] + "/pkgs")

			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else if len(os.Args) == 3 {

		action := os.Args[1]

		if action == "export" {

			// pacman-backup export /mnt/usb-drive
			if isFolder(os.Args[2]) {

				if !isFolder(os.Args[2] + "/sync") {
					makeFolder(os.Args[2] + "/sync")
				}

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				actions.Export(os.Args[2] + "/sync", os.Args[2] + "/pkgs")

			}

		} else if action == "cleanup" {

			// pacman-backup cleanup /mnt/usb-drive
			if isFolder(os.Args[2]) {

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				actions.Cleanup(os.Args[2] + "/pkgs")

			}

		} else if action == "download" {

			// pacman-backup download http://mirror:15678
			if isMirror(os.Args[2]) {

				config := pacman.InitConfig()
				mirror := os.Args[2]

				actions.Sync(mirror, config.Options.DBPath + "/sync", config.Options.CacheDir)
				actions.Download(mirror, config.Options.DBPath + "/sync", config.Options.CacheDir)

			}

		} else if action == "import" {

			// pacman-backup import /mnt/usb-drive
			if isFolder(os.Args[2]) {

				if !isFolder(os.Args[2] + "/sync") {
					makeFolder(os.Args[2] + "/sync")
				}

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				actions.Import(os.Args[2] + "/sync", os.Args[2] + "/pkgs")

			}

		} else if action == "serve" {

			// pacman-backup serve /mnt/usb-drive
			if isFolder(os.Args[2]) {

				if !isFolder(os.Args[2] + "/sync") {
					makeFolder(os.Args[2] + "/sync")
				}

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				actions.Serve(os.Args[2] + "/sync", os.Args[2] + "/pkgs")

			}

		} else if action == "upgrade" {

			// pacman-backup upgrade /mnt/usb-drive
			if isFolder(os.Args[2]) {

				if !isFolder(os.Args[2] + "/sync") {
					makeFolder(os.Args[2] + "/sync")
				}

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				actions.Upgrade(os.Args[2] + "/sync", os.Args[2] + "/pkgs")

			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else if len(os.Args) == 2 {

		action := os.Args[1]

		if action == "cleanup" {

			// pacman-backup cleanup
			config := pacman.InitConfig()

			if isFolder(config.Options.CacheDir) {
				actions.Cleanup(config.Options.CacheDir)
			}

		} else if action == "download" {

			// pacman-backup download
			config := pacman.InitConfig()
			mirror := config.ToMirror()

			if isFolder(config.Options.CacheDir) {
				actions.Sync(mirror, config.Options.DBPath + "/sync", config.Options.CacheDir)
				actions.Download(mirror, config.Options.DBPath + "/sync", config.Options.CacheDir)
			}

		} else if action == "serve" {

			config := pacman.InitConfig()

			if isFolder(config.Options.CacheDir) {
				actions.Serve(config.Options.DBPath + "/sync", config.Options.CacheDir)
			}

		} else if action == "upgrade" {

			config := pacman.InitConfig()

			if isFolder(config.Options.CacheDir) {
				actions.Upgrade(config.Options.DBPath + "/sync", config.Options.CacheDir)
			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else {

		showUsage()
		os.Exit(1)

	}

}
