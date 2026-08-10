package main

import "pacman-backup/actions"
import "pacman-backup/pacman"
import "pacman-backup/structs"
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

func isRootUser() bool {

	user := os.Getenv("USER")

	if user == "root" {
		return true
	}

	return false

}

func showUsage(console *structs.Console) {

	user := os.Getenv("USER")

	if user == "root" {
		user = os.Getenv("SUDO_USER")
	}

	console.Info("")
	console.Info("Pacman Backup")
	console.Info("Offline Pacman Cache Management Tool")
	console.Info("")

	console.Group("Usage: pacman-backup [Action] [Folder]")
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
	console.Log("sudo pacman-backup download;")
	console.Log("pacman-backup export /run/media/" + user + "/pacman-usbdrive;")
	console.Log("pacman-backup cleanup /run/media/" + user + "/pacman-usbdrive;")
	console.Log("sync;")
	console.Log("")
	console.Log("# Step 2: Machine without internet connection")
	console.Log("sudo pacman-backup import /run/media/" + user + "/pacman-usbdrive;")
	console.Log("sudo pacman-backup upgrade /run/media/" + user + "/pacman-usbdrive;")
	console.Log("sync;")
	console.GroupEnd("-----------------")

	console.Group("LAN Mirror Example")
	console.Log("# Step 1: Machine with internet connection")
	console.Log("sudo pacman-backup download;")
	console.Log("sudo pacman-backup serve;")
	console.Log("")
	console.Log("# Step 2: Machine without internet connection")
	console.Log("sudo pacman-backup download http://192.168.0.10:15678")
	console.Log("sudo pacman-backup upgrade;")
	console.Group("------------------")

}

func main() {

	console := structs.NewConsole(os.Stdout, os.Stderr, 0)

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

				actions.Sync(console, mirror, os.Args[3]+"/sync", os.Args[3]+"/pkgs")
				actions.Download(console, mirror, os.Args[3]+"/sync", os.Args[3]+"/pkgs")

			}

		} else {

			showUsage(console)
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

				actions.Export(console, os.Args[2]+"/sync", os.Args[2]+"/pkgs")

			}

		} else if action == "cleanup" {

			// pacman-backup cleanup /mnt/usb-drive
			if isFolder(os.Args[2]) {

				if !isFolder(os.Args[2] + "/sync") {
					makeFolder(os.Args[2] + "/sync")
				}

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				actions.Cleanup(console, os.Args[2]+"/sync", os.Args[2]+"/pkgs")

			}

		} else if action == "download" {

			// pacman-backup download http://mirror:15678
			if isMirror(os.Args[2]) {

				config := pacman.InitConfig("/etc/pacman.conf")
				mirror := os.Args[2]

				if isRootUser() {
					actions.Sync(console, mirror, config.Options.DBPath+"/sync", config.Options.CacheDir)
					actions.Download(console, mirror, config.Options.DBPath+"/sync", config.Options.CacheDir)
				} else {
					console.Error("Please execute this command as the root user")
				}

			// pacman-backup download /mnt/usb-drive
			} else if isFolder(os.Args[2]) {

				config := pacman.InitConfig("/etc/pacman.conf")
				mirror := config.ToMirror()

				if !isFolder(os.Args[2] + "/sync") {
					makeFolder(os.Args[2] + "/sync")
				}

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				actions.Sync(console, mirror, os.Args[2]+"/sync", os.Args[2]+"/pkgs")
				actions.Download(console, mirror, os.Args[2]+"/sync", os.Args[2]+"/pkgs")

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

				if isRootUser() {
					actions.Import(console, os.Args[2]+"/sync", os.Args[2]+"/pkgs")
				} else {
					console.Error("Please execute this command as the root user")
				}

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

				actions.Serve(console, os.Args[2]+"/sync", os.Args[2]+"/pkgs")

			}

		} else if action == "upgrade" {

			// pacman-backup upgrade /mnt/usb-drive
			if isFolder(os.Args[2]) {

				config := pacman.InitConfig("/etc/pacman.conf")
				mirror := config.ToMirror()

				if !isFolder(os.Args[2] + "/sync") {
					makeFolder(os.Args[2] + "/sync")
				}

				if !isFolder(os.Args[2] + "/pkgs") {
					makeFolder(os.Args[2] + "/pkgs")
				}

				if isRootUser() {
					actions.Upgrade(console, mirror, os.Args[2]+"/sync", os.Args[2]+"/pkgs")
				} else {
					console.Error("Please execute this command as the root user")
				}

			}

		} else {

			showUsage(console)
			os.Exit(1)

		}

	} else if len(os.Args) == 2 {

		action := os.Args[1]

		if action == "cleanup" {

			// pacman-backup cleanup
			config := pacman.InitConfig("/etc/pacman.conf")

			if isFolder(config.Options.DBPath+"/sync") && isFolder(config.Options.CacheDir) {

				if isRootUser() {
					actions.Cleanup(console, config.Options.DBPath+"/sync", config.Options.CacheDir)
				} else {
					console.Error("Please execute this command as the root user")
				}

			}

		} else if action == "download" {

			// pacman-backup download
			config := pacman.InitConfig("/etc/pacman.conf")
			mirror := config.ToMirror()

			console.Log(mirror)
			console.Log(config.Options.DBPath)
			console.Log(config.Options.CacheDir)

			if isFolder(config.Options.CacheDir) {

				if isRootUser() {
					actions.Sync(console, mirror, config.Options.DBPath+"/sync", config.Options.CacheDir)
					actions.Download(console, mirror, config.Options.DBPath+"/sync", config.Options.CacheDir)
				} else {
					console.Error("Please execute this command as the root user")
				}

			}

		} else if action == "serve" {

			config := pacman.InitConfig("/etc/pacman.conf")

			if isFolder(config.Options.CacheDir) {
				actions.Serve(console, config.Options.DBPath+"/sync", config.Options.CacheDir)
			}

		} else if action == "upgrade" {

			config := pacman.InitConfig("/etc/pacman.conf")
			mirror := config.ToMirror()

			if isFolder(config.Options.CacheDir) {

				if isRootUser() {
					actions.Upgrade(console, mirror, config.Options.DBPath+"/sync", config.Options.CacheDir)
				} else {
					console.Error("Please execute this command as the root user")
				}

			}

		} else {

			showUsage(console)
			os.Exit(1)

		}

	} else {

		showUsage(console)
		os.Exit(1)

	}

}
