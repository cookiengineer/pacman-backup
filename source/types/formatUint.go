package types

import "strconv"
import "strings"

func formatUint(number uint, length int) string {

	var result string

	chunk := strconv.FormatUint(uint64(number), 10)

	if length == 0 {

		result = chunk

	} else {

		if len(chunk) < length {

			var prefix strings.Builder

			for p := 0; p < length-len(chunk); p++ {
				prefix.WriteString("0")
			}

			result = prefix.String() + chunk

		} else {

			result = chunk

		}

	}

	return result

}
