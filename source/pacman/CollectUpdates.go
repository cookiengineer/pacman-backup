package pacman

import "pacman-backup/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdates(config string) []structs.Package {

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	cmd := exec.Command("pacman", "-Qui", "--noconfirm", "--config", config)
	buffer, err := cmd.Output()

	result := make([]structs.Package, 0)

	if err == nil {

		blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer)), "\n\nName")

		for b := 0; b < len(blocks); b++ {

			block := strings.TrimSpace(blocks[b])

			if block != "" {

				update := structs.NewPackage("pacman")
				ParsePackage("Name "+block, &update)

				if update.Name != "" && update.Version.IsValid() {
					result = append(result, update)
				}

			}

		}

	}

	return result

}
