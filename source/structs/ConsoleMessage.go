package structs

import "runtime"
import "slices"
import "strings"
import "time"

func toConsoleMessageWords(message string) []string {

	result := make([]string, 0)
	chunk  := ""

	for m := 0; m < len(message); m++ {

		chr := string(message[m])

		if chr == " " {

			if chunk != "" {
				result = append(result, chunk)
			}

			chunk = ""

		} else {
			chunk += chr
		}

	}

	if chunk != "" {
		result = append(result, chunk)
	}

	return result

}

type ConsoleMessage struct {
	Time     time.Time `json:"datetime"`
	Method   string    `json:"method"`
	Value    string    `json:"value"`
	Caller   struct {
		File string `json:"file"`
		Line int    `json:"line"`
	} `json:"caller"`
}

func NewConsoleMessage(method string, value string) ConsoleMessage {

	var message ConsoleMessage

	value = strings.ReplaceAll(value, "\t", "    ")

	message.Time = time.Now()
	message.Method = method
	message.Value = value

	// skip NewMessage()
	// skip console.<Method>()
	_, file, line, ok := runtime.Caller(2)

	if ok == true {

		// XXX: Currently there's no way to get the source code / symbols file path
		if strings.Contains(file, "/Software/tholian-network/detective/") {
			file = file[strings.Index(file, "/Software/tholian-network/detective/")+36:]
		}

		if strings.Contains(file, "/Software/tholian-network/endpoint/") {
			file = file[strings.Index(file, "/Software/tholian-network/endpoint/")+35:]
		}

		message.Caller.File = file
		message.Caller.Line = line

	} else {

		message.Caller.File = "???"
		message.Caller.Line = 0

	}

	return message

}

func (message *ConsoleMessage) IsSame(other ConsoleMessage) bool {

	var result bool

	if message.Method == "Progress" && other.Method == "Progress" {

		if strings.Contains(message.Value, " of ") && strings.Contains(other.Value, " of ") {

			words1 := toConsoleMessageWords(message.Value)
			words2 := toConsoleMessageWords(other.Value)

			index1 := slices.Index(words1, "of")
			index2 := slices.Index(words2, "of")

			if index1 > 0 && index1 < len(words1) - 1 && index2 > 0 && index2 < len(words2) - 1 && index1 == index2 && len(words1) == len(words2) {

				prefix1 := strings.Join(words1[0:index1 - 1], " ")
				suffix1 := strings.Join(words1[index1 + 2:], " ")
				prefix2 := strings.Join(words2[0:index2 - 1], " ")
				suffix2 := strings.Join(words2[index2 + 2:], " ")

				// message: "Parsing Incidents (1 of 5) and some other things"
				// other:   "Parsing Incidents (2 of 5) and some other things"
				if prefix1 == prefix2 && suffix1 == suffix2 {
					result = true
				}

			} else if message.Value == other.Value {
				result = true
			}

		} else if strings.Contains(message.Value, " of ") && !strings.Contains(other.Value, " of ") {

			words1 := toConsoleMessageWords(message.Value)
			index1 := slices.Index(words1, "of")

			if index1 > 0 && index1 < len(words1) - 1 {

				prefix1 := strings.Join(words1[0:index1 - 1], " ")
				suffix1 := strings.Join(words1[index1 + 2:], " ")

				// message: "Parsing 1 of 5 Incidents"
				// other:   "Parsing Incidents"
				if strings.HasPrefix(other.Value, prefix1) && strings.HasSuffix(other.Value, suffix1) {
					result = true
				}

			}

		} else if !strings.Contains(message.Value, " of ") && strings.Contains(other.Value, " of ") {

			words2 := toConsoleMessageWords(message.Value)
			index2 := slices.Index(words2, "of")

			if index2 > 0 && index2 < len(words2) - 1 {

				prefix2 := strings.Join(words2[0:index2 - 1], " ")
				suffix2 := strings.Join(words2[index2 + 2:], " ")

				// message: "Parsing Incidents"
				// other:   "Parsing 1 of 5 Incidents"
				if strings.HasPrefix(message.Value, prefix2) && strings.HasSuffix(message.Value, suffix2) {
					result = true
				}

			}

		} else if strings.HasPrefix(message.Value, "Export") && strings.HasPrefix(other.Value, "Export") {
			result = true
		} else if strings.HasPrefix(message.Value, "Parse") && strings.HasPrefix(other.Value, "Parse") {
			result = true
		} else if strings.HasPrefix(message.Value, "Update") && strings.HasPrefix(other.Value, "Update") {
			result = true
		} else if message.Value == other.Value {
			result = true
		}

	} else {

		if message.Method == other.Method && message.Value == other.Value {
			result = true
		}

	}

	return result

}

func (message *ConsoleMessage) Lines() []string {

	result := make([]string, 0)

	if strings.Contains(message.Value, "\n") {

		tmp := strings.Split(message.Value, "\n")

		for _, line := range tmp {

			separator := ""

			if strings.HasPrefix(line, ">") {
				separator = "-"
			} else if strings.HasPrefix(line, "-") {
				separator = "-"
			} else {
				separator = " "
			}

			result = append(result, separator + line)

		}

	} else {

		tmp := []string{message.Value}

		for _, line := range tmp {

			separator := ""

			if strings.HasPrefix(line, ">") {
				separator = "-"
			} else if strings.HasPrefix(line, "-") {
				separator = "-"
			} else {
				separator = " "
			}

			result = append(result, separator + line)

		}

	}

	return result

}
