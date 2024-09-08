package console

import "encoding/json"
import "os"
import "reflect"
import "strings"

var palette map[string]string = map[string]string{
	"default": "\u001b[39m",
	"keyword": "\u001b[38;5;204m",
	"literal": "\u001b[38;5;174m",
	"bool":    "\u001b[38;5;38m",
	"string":  "\u001b[38;5;77m",
	"number":  "\u001b[38;5;197m",
}

func isNumber(chunk string) bool {

	var result bool = true

	for c := 0; c < len(chunk); c++ {

		var character = string(chunk[c])

		if character >= "0" && character <= "9" {
			continue
		} else if character == "+" || character == "-" || character == "." {
			continue
		} else if character == "E" || character == "e" {
			continue
		} else {
			result = false
			break
		}

	}

	return result

}

func toType(instance any) string {

	typ := reflect.TypeOf(instance)

	if typ.Kind() == reflect.Ptr {
		return "*" + typ.Elem().Name()
	} else {
		return typ.Name()
	}

}

func highlight(line string) string {

	var result string
	var suffix string
	var prefix string

	if strings.HasPrefix(line, "    ") {
		prefix = line[0:strings.Index(line, strings.TrimSpace(line))]
		line = line[len(prefix):]
	}

	if strings.HasSuffix(line, ",") {
		suffix = palette["default"] + ","
		line = line[0 : len(line)-1]
	} else {
		suffix = palette["default"]
	}

	if strings.Contains(line, ": ") {

		key := strings.Split(line, ": ")[0]
		val := strings.Join(strings.Split(line, ": ")[1:], ": ")

		if strings.HasPrefix(key, "\"") && strings.HasSuffix(key, "\"") {
			key = palette["string"] + key
		}

		if val == "true" || val == "false" {
			result = key + palette["default"] + ": " + palette["bool"] + val
		} else if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
			result = key + palette["default"] + ": " + palette["string"] + val
		} else if isNumber(val) {
			result = key + palette["default"] + ": " + palette["number"] + val
		} else if val == "[" || val == "{" {
			result = key + palette["default"] + ": " + palette["literal"] + val
		} else if val == "[]" || val == "{}" {
			result = key + palette["default"] + ": " + palette["literal"] + val
		} else if val == "null" || val == "undefined" {
			result = key + palette["default"] + ": " + palette["keyword"] + val
		} else {
			result = key + palette["default"] + ": " + palette["default"] + val
		}

	} else {

		val := line

		if val == "true" || val == "false" {
			result = palette["bool"] + val
		} else if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
			result = palette["string"] + val
		} else if isNumber(val) {
			result = palette["number"] + val
		} else if val == "[" || val == "{" {
			result = palette["literal"] + val
		} else if val == "]" || val == "}" {
			result = palette["literal"] + val
		} else if val == "[]" || val == "{}" {
			result = palette["literal"] + val
		} else if val == "null" || val == "undefined" {
			result = palette["keyword"] + val
		} else {
			result = palette["default"] + val
		}

	}

	return prefix + result + suffix

}

func Inspect(instance any) {

	if features[FeatureInspect] == true {

		var typ string
		var message string

		switch instance.(type) {
		case string:

			tmp := any(instance).(string)

			if strings.HasPrefix(tmp, "{") && strings.HasSuffix(tmp, "}") {
				typ = "JSON"
				message = sanitize(tmp)
			} else if strings.HasPrefix(tmp, "[") && strings.HasSuffix(tmp, "]") {
				typ = "JSON"
				message = sanitize(tmp)
			} else {
				typ = "string"
				message = sanitize(tmp)
			}

		default:

			typ = toType(instance)
			buffer, err := json.MarshalIndent(instance, "", "\t")

			if err == nil {
				message = sanitize(string(buffer))
			}

		}

		offset := toOffset()

		if message != "" {

			MESSAGES = append(MESSAGES, NewMessage("Log", message))

			if COLORS == true {
				os.Stdout.WriteString("\u001b[40m" + offset + " Inspect(" + typ + "):\u001b[K\n")
			} else {
				os.Stdout.WriteString(offset + " Inspect(" + typ + "):\n")
			}

			if strings.Contains(message, "\n") {

				lines := strings.Split(message, "\n")

				if COLORS == true {

					for l := 0; l < len(lines); l++ {
						os.Stdout.WriteString("\u001b[40m" + offset + " " + highlight(lines[l]) + "\u001b[K\n")
					}

					os.Stdout.WriteString("\u001b[0m")

				} else {

					for l := 0; l < len(lines); l++ {
						os.Stdout.WriteString(offset + " " + lines[l] + "\n")
					}

				}

			} else {

				if COLORS == true {
					os.Stdout.WriteString("\u001b[40m" + offset + " " + highlight(message) + "\u001b[K\n")
				} else {
					os.Stdout.WriteString(offset + " " + message + "\n")
				}

			}

		}

	}

}
