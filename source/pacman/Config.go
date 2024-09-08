package pacman

import "os"
import "strings"

type Config struct {
	Options struct {
		DBPath       string   `json:"DBPath"`
		CacheDir     string   `json:"CacheDir"`
		LogFile      string   `json:"LogFile"`
		HoldPkg      []string `json:"HoldPkg"`
		IgnorePkg    []string `json:"IgnorePkg"`
		Architecture string   `json:"Architecture"`
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

	config.Options.DBPath = "/var/lib/pacman"
	config.Options.CacheDir = "/var/cache/pacman/pkg"
	config.Options.LogFile = "/var/log/pacman.log"
	config.Options.HoldPkg = make([]string, 0)
	config.Options.IgnorePkg = make([]string, 0)
	config.Options.Architecture = "auto"

	config.Repositories.Core = make([]string, 0)
	config.Repositories.Extra = make([]string, 0)
	config.Repositories.Community = make([]string, 0)
	config.Repositories.Multilib = make([]string, 0)

	config.Parse("/etc/pacman.conf")

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

						if key == "DBPath" {

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

						} else if key == "Architecture" {
							config.Options.Architecture = val
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

func (config *Config) ToMirror() string {

	var result string

	// TODO: Support maybe a repository parameter?
	if len(config.Repositories.Core) > 0 {
		result = config.Repositories.Core[0]
	}

	return result

}
