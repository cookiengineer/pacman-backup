package types

func isNumber(value string) bool {

	result := true

	for v := 0; v < len(value); v++ {

		chr := string(value[v])

		if chr >= "0" && chr <= "9" {
			continue
		} else {
			result = false
			break
		}

	}

	return result

}
