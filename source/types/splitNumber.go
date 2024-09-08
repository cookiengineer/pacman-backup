package types

func splitNumber(value string) (string, string, string) {

	var prefix string
	var number string
	var suffix string

	mode := "prefix"

	for v := 0; v < len(value); v++ {

		character := string(value[v])

		if mode == "prefix" {

			if character >= "0" && character <= "9" {
				number += character
				mode = "number"
			} else if character != " " {
				prefix += character
			}

		} else if mode == "number" {

			if character >= "0" && character <= "9" {
				number += character
			} else if character != " " {
				suffix += character
				mode = "suffix"
			}

		} else if mode == "suffix" {

			if character != " " {
				suffix += character
			}

		}

	}

	return prefix, number, suffix

}
