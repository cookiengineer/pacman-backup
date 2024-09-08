package pacman

import "pacman-backup/console"
import "pacman-backup/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdates(config string) []structs.Package {

	update_index := make(map[string]bool, 0)

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	cmd1 := exec.Command("pacman", "-Qu", "--noconfirm", "--config", config)
	buffer1, err1 := cmd1.Output()

	if err1 == nil {

		lines := strings.Split(strings.TrimSpace(string(buffer1)), "\n")

		for l := 0; l < len(lines); l++ {

			line := strings.TrimSpace(lines[l])

			if strings.HasSuffix(line, "[ignored]") {
				line = strings.TrimSpace(line[0 : len(line)-9])
			}

			if strings.Contains(line, " ") && strings.Contains(line, " -> ") {

				// "package 1.2.3 -> 1.2.4"

				name := line[0:strings.Index(line, " ")]
				update_index[name] = true

			}

		}

	} else {
		console.Error(err1.Error())
	}

	cmd := exec.Command("pacman", "-Si", "--noconfirm", "--config", config)
	buffer, err2 := cmd.Output()

	result := make([]structs.Package, 0)

	if err2 == nil {

		blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer)), "\n\nRepository")

		for b := 0; b < len(blocks); b++ {

			block := strings.TrimSpace(blocks[b])

			if block != "" {

				update := structs.NewPackage("pacman")
				ParsePackage("Repository "+block, &update)

				if update.Name != "" && update.Version.IsValid() {

					_, ok := update_index[update.Name]

					if ok == true {
						result = append(result, update)
					}

				}

			}

		}

	} else {
		console.Error(err2.Error())
	}

	return result

}
