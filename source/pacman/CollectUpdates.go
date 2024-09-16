package pacman

import "pacman-backup/console"
import "pacman-backup/structs"
import "bytes"
import "os"
import "os/exec"
import "strings"

func CollectUpdates(config string) []structs.Package {

	update_index := make(map[string]bool, 0)
	result := make([]structs.Package, 0)

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd1 := exec.Command("pacman", "-Qu", "--noconfirm", "--config", config)
	cmd1.Stdout = &stdout
	cmd1.Stderr = &stderr
	err1 := cmd1.Run()

	if err1 == nil {

		lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")

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

		cmd2 := exec.Command("pacman", "-Si", "--noconfirm", "--config", config)
		buffer, err2 := cmd2.Output()

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

	} else {
		console.Error(stderr.String())
	}

	return result

}
