package console

import "os"
import "strings"

var COLORS bool = true
var MESSAGES []Message
var OFFSET int = 0

func init() {

	MESSAGES = make([]Message, 0)

	term := strings.ToLower(os.Getenv("TERM"))

	if term == "xterm" {
		COLORS = false
	} else if term == "xterm-16color" {
		COLORS = false
	} else if term == "xterm-88color" {
		COLORS = true
	} else if term == "xterm-256color" {
		COLORS = true
	} else if term == "xterm-kitty" {
		COLORS = true
	}

	no_color := strings.ToLower(os.Getenv("NO_COLOR"))

	if no_color != "" {

		if no_color == "yes" || no_color == "true" || no_color == "1" {
			COLORS = false
		}

	}

}

func sanitize(message string) string {

	var result string = message

	result = strings.ReplaceAll(result, "\t", "    ")

	return result

}

func toOffset() string {

	var offset string

	if OFFSET > 0 {

		offset = "|"

		for o := 1; o < OFFSET; o++ {
			offset += "|"
		}

	}

	return offset

}

func toSeparator(message string) string {

	var separator string

	if strings.HasPrefix(message, ">") {
		separator = "-"
	} else if strings.HasPrefix(message, "-") {
		separator = "-"
	} else {
		separator = " "
	}

	return separator

}
