package matchers

import "pacman-backup/types"
import "strings"

func parseVersionCondition(value string) (string, string, string) {

	var result_name string
	var result_version string
	var result_architecture string

	if strings.Contains(value, ".c32") {

		if strings.HasSuffix(value, "=0") {

			// "libcom32.c32=0"
			result_name = strings.TrimSpace(value[0 : len(value)-2])
			result_version = "any"

		} else if strings.HasSuffix(value, "=0.0.0") {

			// "libcom32.c32=0.0.0"
			result_name = strings.TrimSpace(value[0 : len(value)-6])
			result_version = "any"

		}

	} else if strings.Contains(value, ".so") {

		if strings.HasSuffix(value, "=0") {

			// "libwhatever.so=0"
			result_name = strings.TrimSpace(value[0 : len(value)-2])
			result_version = "any"

		} else if strings.HasSuffix(value, "=0.0.0") {

			// "libwhatever.so=0.0.0"
			result_name = strings.TrimSpace(value[0 : len(value)-6])
			result_version = "any"

		} else if strings.Contains(value, "=") {

			result_name = strings.TrimSpace(value[0:strings.Index(value, "=")])

			tmp_version := strings.TrimSpace(value[strings.Index(value, "=")+1:])

			if strings.HasSuffix(tmp_version, "-32") {

				// "libwhatever.so=1.2.3-32"
				result_version = "= " + tmp_version[0:len(tmp_version)-3]
				result_architecture = "x86"

			} else if strings.HasSuffix(tmp_version, "-64") {

				// "libwhatever.so=1.2.3-64"
				result_version = "= " + tmp_version[0:len(tmp_version)-3]
				result_architecture = "x86_64"

			} else {

				version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, "=")+1:]))

				result_name = strings.TrimSpace(value[0:strings.Index(value, "=")])
				result_version = "= " + version.String()

			}

		} else if strings.Contains(value, " ") {

			result_name = strings.TrimSpace(value[0:strings.Index(value, " ")])

			tmp := strings.TrimSpace(value[strings.Index(value, " ")+1:])

			if tmp != "" {

				version := types.ToVersion(tmp)
				result_version = "= " + version.String()

			}

		}

	} else if strings.Contains(value, " (") && strings.HasSuffix(value, ")") {

		// "package (1.2.3)"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, " (")+2 : len(value)-1]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, " (")])
		result_version = "= " + version.String()

	} else if strings.HasSuffix(value, ":any") {

		// "package:any"

		result_name = strings.TrimSpace(value[0 : len(value)-4])
		result_version = "any"

	} else if strings.Contains(value, ">=") {

		// "package >= 1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, ">=")+2:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, ">=")])
		result_version = ">= " + version.String()

	} else if strings.Contains(value, ">>") {

		// "package >> 1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, ">>")+2:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, ">>")])
		result_version = "> " + version.String()

	} else if strings.Contains(value, "<=") && strings.HasSuffix(value, "^") {

		// "package <= 1.2.3+hash+commit^"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, "<=")+2 : len(value)-1]))
		version = version.ToLatest(true, true)

		result_name = strings.TrimSpace(value[0:strings.Index(value, "<=")])
		result_version = "<= " + version.String()

	} else if strings.Contains(value, "<=") {

		// "package <= 1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, "<=")+2:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, "<=")])
		result_version = "<= " + version.String()

	} else if strings.Contains(value, "<<") {

		// "package << 1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, "<<")+2:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, "<<")])
		result_version = "< " + version.String()

	} else if strings.Contains(value, ">") {

		// "package > 1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, ">")+1:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, ">")])
		result_version = "> " + version.String()

	} else if strings.Contains(value, "<") {

		// "package < 1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, "<")+1:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, "<")])
		result_version = "< " + version.String()

	} else if strings.Contains(value, "=") {

		// "package = 1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, "=")+1:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, "=")])
		result_version = "= " + version.String()

	} else if strings.Contains(value, " ^") {

		// "package ^1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, " ^")+2:]))
		version = version.ToLatest(true, true)

		result_name = strings.TrimSpace(value[0:strings.Index(value, " ^")])
		result_version = "<= " + version.String()

	} else if strings.Contains(value, " ~") {

		// "package ~1.2.3"

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, " ~")+2:]))
		version = version.ToLatest(false, true)

		result_name = strings.TrimSpace(value[0:strings.Index(value, " ~")])
		result_version = "<= " + version.String()

	} else if strings.Contains(value, " ") {

		version := types.ToVersion(strings.TrimSpace(value[strings.Index(value, " ")+1:]))

		result_name = strings.TrimSpace(value[0:strings.Index(value, " ")])
		result_version = "= " + version.String()

	} else {

		result_name = value
		result_version = "any"

	}

	// Dirty Hack?
	if strings.HasPrefix(result_name, "lib32-") {
		result_architecture = "x86"
	}

	return result_name, result_version, result_architecture

}
