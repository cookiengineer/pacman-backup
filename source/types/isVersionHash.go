package types

func isVersionHash(value string) bool {

	var result bool = false

	// git-describe prefixes hashes with "g"
	if len(value) > 1 && string(value[0]) == "g" {
		value = value[1:]
	}

	if len(value) >= 6 {

		result = true

		for v := 0; v < len(value); v++ {

			character := string(value[v])

			if character >= "0" && character <= "9" {
				continue
			} else if character >= "a" && character <= "f" {
				continue
			} else {
				result = false
				break
			}

		}

	}

	return result

}
