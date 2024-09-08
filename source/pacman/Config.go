package pacman

import "os"
import "strconv"
import "strings"

type Config struct {
	Options struct {
		RootDir            string   `json:"RootDir"`
		DBPath             string   `json:"DBPath"`
		CacheDir           string   `json:"CacheDir"`
		LogFile            string   `json:"LogFile"`
		// GPGDir = /etc/pacman.d/gnupg/
		// HookDir = /etc/pacman.d/hooks/
		HoldPkg            []string `json:"HoldPkg"`
		IgnorePkg          []string `json:"IgnorePkg"`
		IgnoreGroup        []string `json:"IgnoreGroup"`
		SyncFirst          []string `json:"SyncFirst"`

		CleanMethod        string   `json:"CleanMethod"`
		Architecture       string   `json:"Architecture"`
		NoUpgrade          []string `json:"NoUpgrade"`
		NoExtract          []string `json:"NoExtract"`

		UseSyslog          bool     `json:"UseSyslog"`
		// Color
		NoProgressBar      bool     `json:"NoProgressBar"`
		CheckSpace         bool     `json:"CheckSpace"`
		// VerbosePkgLists
		ParallelDownloads  int      `json:"ParallelDownloads"`

		SigLevel           string   `json:"SigLevel"`
		LocalFileSigLevel  string   `json:"LocalFileSigLevel"`
		RemoteFileSigLevel string   `json:"RemoteFileSigLevel"`
	} `json:"options"`
	Repositories struct {
		Core      []string `json:"core"`
		Extra     []string `json:"extra"`
		Community []string `json:"community"`
		Multilib  []string `json:"multilib"`
	} `json:"repositories"`
}

func InitConfig() Config {

	var config Config

	config.Options.RootDir = "/"
	config.Options.DBPath = "/var/lib/pacman"
	config.Options.CacheDir = "/var/cache/pacman/pkg"
	config.Options.LogFile = "/var/log/pacman.log"
	config.Options.HoldPkg = make([]string, 0)
	config.Options.IgnorePkg = make([]string, 0)
	config.Options.IgnoreGroup = make([]string, 0)
	config.Options.SyncFirst = make([]string, 0)

	config.Options.CleanMethod = "KeepInstalled"
	config.Options.Architecture = "auto"
	config.Options.NoUpgrade = make([]string, 0)
	config.Options.NoExtract = make([]string, 0)

	config.Options.UseSyslog = false
	config.Options.NoProgressBar = false
	config.Options.CheckSpace = true
	config.Options.ParallelDownloads = 1

	config.Options.SigLevel = "Required DatabaseOptional"
	config.Options.LocalFileSigLevel = "Optional"
	config.Options.RemoteFileSigLevel = "Required"

	config.Repositories.Core = make([]string, 0)
	config.Repositories.Extra = make([]string, 0)
	config.Repositories.Community = make([]string, 0)
	config.Repositories.Multilib = make([]string, 0)

	config.Parse("/etc/pacman.conf")

	return config

}

func NewConfig(mirror string, sync_folder string, pkgs_folder string) Config {

	var config Config

	config.Options.RootDir = "/"

	if strings.HasSuffix(sync_folder, "/sync") {
		config.Options.DBPath = sync_folder[0:len(sync_folder)-5]
	} else {
		config.Options.DBPath = sync_folder
	}

	config.Options.CacheDir = pkgs_folder

	config.Options.LogFile = "/var/log/pacman.log"
	config.Options.HoldPkg = []string{"pacman", "glibc"}
	config.Options.IgnorePkg = []string{"linux", "linux-headers", "linux-firmware"}
	config.Options.IgnoreGroup = make([]string, 0)
	config.Options.SyncFirst = make([]string, 0)

	config.Options.CleanMethod = "KeepInstalled"
	config.Options.Architecture = "auto"
	config.Options.NoUpgrade = make([]string, 0)
	config.Options.NoExtract = make([]string, 0)

	config.Options.UseSyslog = false
	config.Options.NoProgressBar = false
	config.Options.CheckSpace = true
	config.Options.ParallelDownloads = 1

	config.Options.SigLevel = "Optional DatabaseOptional"
	config.Options.LocalFileSigLevel = "Optional"
	config.Options.RemoteFileSigLevel = "Optional"

	config.Repositories.Core = make([]string, 0)
	config.Repositories.Extra = make([]string, 0)
	config.Repositories.Community = make([]string, 0)
	config.Repositories.Multilib = make([]string, 0)

	if strings.HasPrefix(mirror, "https://") || strings.HasPrefix(mirror, "http://") {

		config.Repositories.Core = append(config.Repositories.Core, mirror)
		config.Repositories.Extra = append(config.Repositories.Extra, mirror)
		config.Repositories.Community = append(config.Repositories.Community, mirror)
		config.Repositories.Multilib = append(config.Repositories.Multilib, mirror)

	}

	return config

}

func (config *Config) Parse(file string) {

	buffer, err := os.ReadFile(file)

	if err == nil {

		lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")
		section := ""

		for l := 0; l < len(lines); l++ {

			line := strings.TrimSpace(lines[l])

			if strings.Contains(line, "#") {
				line = strings.TrimSpace(line[0:strings.Index(line, "#")])
			}

			if line != "" {

				if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {

					section = strings.TrimSpace(line[1:len(line)-1])

				} else if strings.Contains(line, " = ") {

					key := strings.TrimSpace(line[0:strings.Index(line, " = ")])
					val := strings.TrimSpace(line[strings.Index(line, " = ")+3:])

					if section == "options" {

						if key == "RootDir" {

							if strings.HasPrefix(val, "/") && strings.HasSuffix(val, "/") {
								config.Options.RootDir = val[0:len(val)-1]
							} else if strings.HasPrefix(val, "/") {
								config.Options.RootDir = val
							}

						} else if key == "DBPath" {

							if strings.HasPrefix(val, "/") && strings.HasSuffix(val, "/") {
								config.Options.DBPath = val[0:len(val)-1]
							} else if strings.HasPrefix(val, "/") {
								config.Options.DBPath = val
							}

						} else if key == "CacheDir" {

							if strings.HasPrefix(val, "/") && strings.HasSuffix(val, "/") {
								config.Options.CacheDir = val[0:len(val)-1]
							} else if strings.HasPrefix(val, "/") {
								config.Options.CacheDir = val
							}

						} else if key == "LogFile" {

							if strings.HasPrefix(val, "/") {
								config.Options.LogFile = val
							}

						} else if key == "HoldPkg" {

							tmp := strings.Split(val, " ")

							for t := 0; t < len(tmp); t++ {
								config.Options.HoldPkg = append(config.Options.HoldPkg, tmp[t])
							}

						} else if key == "IgnorePkg" {

							tmp := strings.Split(val, " ")

							for t := 0; t < len(tmp); t++ {
								config.Options.IgnorePkg = append(config.Options.IgnorePkg, tmp[t])
							}

						} else if key == "IgnoreGroup" {

							tmp := strings.Split(val, " ")

							for t := 0; t < len(tmp); t++ {
								config.Options.IgnoreGroup = append(config.Options.IgnoreGroup, tmp[t])
							}

						} else if key == "SyncFirst" {

							tmp := strings.Split(val, " ")

							for t := 0; t < len(tmp); t++ {
								config.Options.SyncFirst = append(config.Options.SyncFirst, tmp[t])
							}

						} else if key == "CleanMethod" {

							config.Options.CleanMethod = val

						} else if key == "Architecture" {

							config.Options.Architecture = val

						} else if key == "NoUpgrade" {

							tmp := strings.Split(val, " ")

							for t := 0; t < len(tmp); t++ {
								config.Options.NoUpgrade = append(config.Options.NoUpgrade, tmp[t])
							}

						} else if key == "NoExtract" {

							tmp := strings.Split(val, " ")

							for t := 0; t < len(tmp); t++ {
								config.Options.NoExtract = append(config.Options.NoExtract, tmp[t])
							}

						} else if key == "ParallelDownloads" {

							num, err := strconv.ParseInt(val, 10, 16)

							if err == nil {
								config.Options.ParallelDownloads = int(num)
							}

						} else if key == "SigLevel" {

							config.Options.SigLevel = val

						} else if key == "LocalFileSigLevel" {

							config.Options.LocalFileSigLevel = val

						} else if key == "RemoteFileSigLevel" {

							config.Options.RemoteFileSigLevel = val

						}

					} else if section == "core" || section == "extra" || section == "community" || section == "multilib" {

						if key == "Include" {

							config.ParseMirrorlist(section, val)

						} else if key == "Server" {

							if strings.HasPrefix(val, "http://") || strings.HasPrefix(val, "https://") {
							
								if section == "core" {
									config.Repositories.Core = append(config.Repositories.Core, val)
								} else if section == "extra" {
									config.Repositories.Extra = append(config.Repositories.Extra, val)
								} else if section == "community" {
									config.Repositories.Community = append(config.Repositories.Community, val)
								} else if section == "multilib" {
									config.Repositories.Multilib = append(config.Repositories.Multilib, val)
								}

							}

						}

					}

				} else {

					if section == "options" {

						key := strings.TrimSpace(line)

						if key == "UseSysLog" {
							config.Options.UseSyslog = true
						} else if key == "NoProgressBar" {
							config.Options.NoProgressBar = true
						} else if key == "CheckSpace" {
							config.Options.CheckSpace = true
						}

					}

				}

			}

		}

	}

}

func (config *Config) ParseMirrorlist(repository string, file string) {

	buffer, err := os.ReadFile(file)

	if err == nil {

		lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

		for l := 0; l < len(lines); l++ {

			line := strings.TrimSpace(lines[l])

			if strings.Contains(line, "#") {
				line = strings.TrimSpace(line[0:strings.Index(line, "#")])
			}

			if line != "" && strings.Contains(line, " = ") {

				key := strings.TrimSpace(line[0:strings.Index(line, " = ")])
				val := strings.TrimSpace(line[strings.Index(line, " = ")+3:])

				if key == "Server" {

					if strings.HasPrefix(val, "http://") || strings.HasPrefix(val, "https://") {
					
						if repository == "core" {
							config.Repositories.Core = append(config.Repositories.Core, val)
						} else if repository == "extra" {
							config.Repositories.Extra = append(config.Repositories.Extra, val)
						} else if repository == "community" {
							config.Repositories.Community = append(config.Repositories.Community, val)
						} else if repository == "multilib" {
							config.Repositories.Multilib = append(config.Repositories.Multilib, val)
						}

					}

				}

			}

		}

	}

}

func (config *Config) String() string {

	lines := make([]string, 0)

	lines = append(lines, "")
	lines = append(lines, "[options]")

	if config.Options.DBPath != "" {
		lines = append(lines, "DBPath = " + config.Options.DBPath + "/")
	}

	if config.Options.CacheDir != "" {
		lines = append(lines, "CacheDir = " + config.Options.CacheDir + "/")
	}

	if config.Options.LogFile != "" {
		lines = append(lines, "LogFile = " + config.Options.LogFile)
	}

	if len(config.Options.HoldPkg) > 0 {
		lines = append(lines, "HoldPkg = " + strings.Join(config.Options.HoldPkg, " "))
	}

	if len(config.Options.IgnorePkg) > 0 {
		lines = append(lines, "IgnorePkg = " + strings.Join(config.Options.IgnorePkg, " "))
	}

	if len(config.Options.IgnoreGroup) > 0 {
		lines = append(lines, "IgnoreGroup = " + strings.Join(config.Options.IgnoreGroup, " "))
	}

	if len(config.Options.SyncFirst) > 0 {
		lines = append(lines, "SyncFirst = " + strings.Join(config.Options.SyncFirst, " "))
	}

	if config.Options.CleanMethod != "" {
		lines = append(lines, "CleanMethod = " + config.Options.CleanMethod)
	}

	if config.Options.Architecture != "" {
		lines = append(lines, "Architecture = " + config.Options.Architecture)
	}

	if len(config.Options.NoUpgrade) > 0 {
		lines = append(lines, "NoUpgrade = " + strings.Join(config.Options.NoUpgrade, " "))
	}

	if len(config.Options.NoExtract) > 0 {
		lines = append(lines, "NoExtract = " + strings.Join(config.Options.NoExtract, " "))
	}

	if config.Options.UseSyslog {
		lines = append(lines, "UseSyslog")
	}

	if config.Options.NoProgressBar {
		lines = append(lines, "NoProgressBar")
	}

	if config.Options.CheckSpace {
		lines = append(lines, "CheckSpace")
	}

	if config.Options.ParallelDownloads > 0 {
		lines = append(lines, "ParallelDownloads = " + strconv.Itoa(config.Options.ParallelDownloads))
	}

	if config.Options.SigLevel != "" {
		lines = append(lines, "SigLevel = " + config.Options.SigLevel)
	}

	if config.Options.LocalFileSigLevel != "" {
		lines = append(lines, "LocalFileSigLevel = " + config.Options.LocalFileSigLevel)
	}

	if config.Options.RemoteFileSigLevel != "" {
		lines = append(lines, "RemoteFileSigLevel = " + config.Options.RemoteFileSigLevel)
	}

	lines = append(lines, "")
	lines = append(lines, "[core]")

	if len(config.Repositories.Core) > 0 {

		for r := 0; r < len(config.Repositories.Core); r++ {
			lines = append(lines, "Server = " + config.Repositories.Core[r])
		}

	}

	lines = append(lines, "")
	lines = append(lines, "[extra]")

	if len(config.Repositories.Extra) > 0 {

		for r := 0; r < len(config.Repositories.Extra); r++ {
			lines = append(lines, "Server = " + config.Repositories.Extra[r])
		}

	}

	lines = append(lines, "")
	lines = append(lines, "[community]")

	if len(config.Repositories.Community) > 0 {

		for r := 0; r < len(config.Repositories.Community); r++ {
			lines = append(lines, "Server = " + config.Repositories.Community[r])
		}

	}

	lines = append(lines, "")
	lines = append(lines, "[multilib]")

	if len(config.Repositories.Multilib) > 0 {

		for r := 0; r < len(config.Repositories.Multilib); r++ {
			lines = append(lines, "Server = " + config.Repositories.Multilib[r])
		}

	}

	return strings.Join(lines, "\n")

}

func (config *Config) ToMirror() string {

	var result string

	// TODO: Support maybe a repository parameter?
	if len(config.Repositories.Core) > 0 {
		result = config.Repositories.Core[0]
	}

	return result

}
