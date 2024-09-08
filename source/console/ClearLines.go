package console

import "os"
import "strconv"

func ClearLines(number int) {

	if number < 32 {

		// move up x lines
		os.Stdout.WriteString("\033[" + strconv.Itoa(number) + "A")

		// clear from cursor to end of screen
		os.Stdout.WriteString("\033[J")

	}

}
