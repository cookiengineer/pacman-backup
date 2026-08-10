package structs

import "golang.org/x/term"
import "os"
import "strings"
import "sync"

func toConsoleIndent(offset int) string {

	result := ""

	if offset > 0 {

		result = "|"

		for o := 1; o < offset; o++ {
			result += "|"
		}

	}

	return result

}

type Console struct {
	Messages   []ConsoleMessage `json:"messages"`
	Stdout     *os.File         `json:"stdout"`
	Stderr     *os.File         `json:"stderr"`
	methods    map[string]bool  `json:"-"`
	mutex      *sync.RWMutex    `json:"-"`
	offset     int              `json:"-"`
	width      int              `json:"-"`
	height     int              `json:"-"`
	use_colors bool             `json:"-"`
}

func NewConsole(stdout *os.File, stderr *os.File, offset int) *Console {

	var console Console

	if offset >= 0 {
		console.offset = offset
	} else {
		console.offset = 0
	}

	console.Messages = make([]ConsoleMessage, 0)
	console.methods = map[string]bool{
		"Group": true,
		"Log": true,
		"Info": true,
		"Warn": true,
		"Error": true,
		"Progress": true,
	}
	console.mutex = &sync.RWMutex{}

	if stdout != nil && stderr != nil {

		if term.IsTerminal(int(stdout.Fd())) == true {

			width, height, err := term.GetSize(int(stdout.Fd()))

			if err == nil {
				console.width = width
				console.height = height
			} else {
				console.width = 0
				console.height = 0
			}

		}

		console.Stdout = stdout
		console.Stderr = stderr

		term := strings.ToLower(os.Getenv("TERM"))

		if term == "xterm" {
			console.use_colors = true
		} else if term == "xterm-16color" {
			console.use_colors = true
		} else if term == "xterm-88color" {
			console.use_colors = true
		} else if term == "xterm-256color" {
			console.use_colors = true
		} else if term == "xterm-kitty" {
			console.use_colors = true
		}

		no_color := strings.ToLower(os.Getenv("NO_COLOR"))

		if no_color == "yes" || no_color == "true" || no_color == "1" {
			console.use_colors = false
		}

	} else {

		console.Stdout = nil
		console.Stderr = nil
		console.use_colors = false
		console.width = 0
		console.height = 0

	}

	return &console

}

func (console *Console) GetOffset() int {

	var result int

	console.mutex.RLock()

	result = console.offset

	console.mutex.RUnlock()

	return result

}

func (console *Console) HasErrors() bool {

	var result bool

	console.mutex.RLock()

	for _, message := range console.Messages {

		if message.Method == "Error" {
			result = true
			break
		}

	}

	console.mutex.RUnlock()

	return result

}

func (console *Console) SetOffset(offset int) bool {

	var result bool

	console.mutex.Lock()

	if offset >= 0 {
		console.offset = offset
		result = true
	}

	console.mutex.Unlock()

	return result

}

func (console *Console) UseColors() bool {

	var result bool

	console.mutex.RLock()

	if console.use_colors == true {
		result = true
	}

	console.mutex.RUnlock()

	return result

}

func (console *Console) UseMethod(method string) bool {

	var result bool

	console.mutex.RLock()

	is_enabled, ok := console.methods[method]

	if ok == true {
		result = is_enabled
	}

	console.mutex.RUnlock()

	return result

}

func (console *Console) DisableMethod(method string) bool {

	var result bool

	console.mutex.Lock()

	_, ok := console.methods[method]

	if ok == true {
		console.methods[method] = false
		result = true
	}

	console.mutex.Unlock()

	return result

}

func (console *Console) EnableMethod(method string) bool {

	var result bool

	console.mutex.Lock()

	_, ok := console.methods[method]

	if ok == true {
		console.methods[method] = true
		result = true
	}

	console.mutex.Unlock()

	return result

}

func (console *Console) Render(target *Console) {

	console.mutex.RLock()

	preserved := console.Messages[0:len(console.Messages)]

	if target.height > 0 && len(preserved) > target.height {

		target.mutex.Lock()
		target.Messages = append(target.Messages, preserved[0:len(preserved) - target.height]...)
		target.mutex.Unlock()

		preserved = preserved[len(preserved) - target.height:]

	}

	for _, message := range preserved {

		if message.Method == "Clear" {
			target.Clear(message.Value)
		} else if message.Method == "Error" {
			target.Error(message.Value)
		} else if message.Method == "Group" {
			target.Group(message.Value)
		} else if message.Method == "GroupEnd" {
			target.GroupEnd(message.Value)
		} else if message.Method == "Info" {
			target.Info(message.Value)
		} else if message.Method == "Log" {
			target.Log(message.Value)
		} else if message.Method == "Progress" {
			target.Progress(message.Value)
		} else if message.Method == "Warn" {
			target.Warn(message.Value)
		}

	}

	console.mutex.RUnlock()

}

func (console *Console) Reset(offset bool, errors bool, warnings bool) {

	console.mutex.Lock()

	if offset == true {
		console.offset = 0
	}

	if errors == false || warnings == false {

		preserved := make([]ConsoleMessage, 0)

		for _, message := range console.Messages {

			if message.Method == "Error" && errors == false {
				preserved = append(preserved, message)
			} else if message.Method == "Warn" && warnings == false {
				preserved = append(preserved, message)
			}

		}

		console.Messages = preserved

	} else {
		console.Messages = make([]ConsoleMessage, 0)
	}

	console.mutex.Unlock()

}

func (console *Console) Length() int {

	var result int

	console.mutex.RLock()

	result = len(console.Messages)

	console.mutex.RUnlock()

	return result

}

// func (console *Console) ClearScreen() {
// 
// 	message := NewConsoleMessage("Clear", "")
// 
// 	if console.HasErrors() == false {
// 
// 		if console.Stdout != nil {
// 
// 			// Clear Screen and Reset Cursor
// 			console.Stdout.WriteString("\u001b[2J\u001b[0f")
// 
// 			// Clear Scroll Buffer
// 			console.Stdout.WriteString("\u001b[3J")
// 
// 		}
// 
// 		console.mutex.Lock()
// 		console.Messages = append(console.Messages, message)
// 		console.mutex.Unlock()
// 
// 	}
// 
// }

func (console *Console) Clear(raw string) {

	// Multiline is forbidden
	if strings.Contains(raw, "\n") {
		raw = raw[0:strings.Index(raw, "\n")]
	}

	found_start := -1
	found_end   := -1

	for m := 0; m < len(console.Messages); m++ {

		message := console.Messages[m]

		if message.Method == "Group" {

			if strings.HasPrefix(message.Value, raw) {
				found_start = m
				break
			}

		}

	}

	if found_start >= 0 {

		for m := found_start; m < len(console.Messages); m++ {

			message := console.Messages[m]

			if message.Method == "GroupEnd" {

				if m > found_start && strings.HasPrefix(message.Value, raw) {
					found_end = m
					break
				}

			}

		}

	}

	if found_start != -1 && found_end == -1 {

		preserved := make([]ConsoleMessage, 0)

		for m := 0; m <= found_start; m++ {
			preserved = append(preserved, console.Messages[m])
		}

		console.mutex.Lock()
		console.Messages = make([]ConsoleMessage, 0)
		console.offset   = 0
		console.mutex.Unlock()

		if console.Stdout != nil {

			// Clear Screen and Reset Cursor
			console.Stdout.WriteString("\u001b[2J\u001b[0f")

			// Clear Scroll Buffer
			console.Stdout.WriteString("\u001b[3J")

		}

		if console.height > 0 && len(preserved) > console.height {

			console.mutex.Lock()
			console.Messages = preserved[0:len(preserved) - console.height]
			console.mutex.Unlock()

			preserved = preserved[len(preserved) - console.height:]

		}

		for _, message := range preserved {

			if message.Method == "Clear" {
				console.Clear(message.Value)
			} else if message.Method == "Error" {
				console.Error(message.Value)
			} else if message.Method == "Group" {
				console.Group(message.Value)
			} else if message.Method == "GroupEnd" {
				console.GroupEnd(message.Value)
			} else if message.Method == "Info" {
				console.Info(message.Value)
			} else if message.Method == "Log" {
				console.Log(message.Value)
			} else if message.Method == "Progress" {
				console.Progress(message.Value)
			} else if message.Method == "Warn" {
				console.Warn(message.Value)
			}

		}

	}

}

func (console *Console) Group(raw string) {

	// Multiline is forbidden
	if strings.Contains(raw, "\n") {
		raw = raw[0:strings.Index(raw, "\n")]
	}

	message := NewConsoleMessage("Group", raw)

	if enabled, _ := console.methods["Group"]; enabled == true {

		indent := toConsoleIndent(console.offset)
		lines := message.Lines()

		if len(lines) == 1 && console.Stdout != nil {

			if console.use_colors == true {
				console.Stdout.WriteString("\u001b[40m" + indent + "/" + lines[0] + "\u001b[K\u001b[0m\n")
			} else {
				console.Stdout.WriteString(indent + "/" + lines[0] + "\n")
			}

		}

		console.mutex.Lock()
		console.Messages = append(console.Messages, message)
		console.offset += 1
		console.mutex.Unlock()

	}

}

func (console *Console) GroupEnd(raw string) {

	// Multiline is forbidden
	if strings.Contains(raw, "\n") {
		raw = raw[0:strings.Index(raw, "\n")]
	}

	message := NewConsoleMessage("GroupEnd", raw)

	if enabled, _ := console.methods["Group"]; enabled == true {

		console.mutex.Lock()

		if console.offset > 0 {
			console.offset -= 1
		}
		console.mutex.Unlock()

		indent := toConsoleIndent(console.offset)
		lines := message.Lines()

		if len(lines) == 1 && console.Stdout != nil {

			if console.use_colors == true {
				console.Stdout.WriteString("\u001b[40m" + indent + "\\" + lines[0] + "\u001b[K\u001b[0m\n")
			} else {
				console.Stdout.WriteString(indent + "\\" + lines[0] + "\n")
			}

		}

		console.mutex.Lock()
		console.Messages = append(console.Messages, message)
		console.mutex.Unlock()

	}

}

func (console *Console) Log(raw string) {

	message := NewConsoleMessage("Log", raw)

	if enabled, _ := console.methods["Log"]; enabled == true {

		indent := toConsoleIndent(console.offset)
		lines  := message.Lines()

		if len(lines) > 0 && console.Stdout != nil {

			if console.use_colors == true {

				for _, line := range lines {
					console.Stdout.WriteString("\u001b[40m" + indent + line + "\u001b[K\n")
				}

				console.Stdout.WriteString("\u001b[0m")

			} else {

				for _, line := range lines {
					console.Stdout.WriteString(indent + line + "\n")
				}

			}

		}

		console.mutex.Lock()
		console.Messages = append(console.Messages, message)
		console.mutex.Unlock()

	}

}

func (console *Console) Error(raw string) {

	message := NewConsoleMessage("Error", raw)

	if enabled, _ := console.methods["Error"]; enabled == true {

		indent := toConsoleIndent(console.offset)
		lines  := message.Lines()

		if len(lines) > 0 && console.Stderr != nil {

			if console.use_colors == true {

				for _, line := range lines {
					console.Stderr.WriteString("\u001b[41m" + indent + line + "\u001b[K\n")
				}

				console.Stderr.WriteString("\u001b[0m")

			} else {

				for _, line := range lines {
					console.Stderr.WriteString(indent + line + "\n")
				}

			}

		}

		console.mutex.Lock()
		console.Messages = append(console.Messages, message)
		console.mutex.Unlock()

	}

}

func (console *Console) Info(raw string) {

	message := NewConsoleMessage("Info", raw)

	if enabled, _ := console.methods["Info"]; enabled == true {

		indent := toConsoleIndent(console.offset)
		lines  := message.Lines()

		if len(lines) > 0 && console.Stdout != nil {

			if console.use_colors == true {

				for _, line := range lines {
					console.Stdout.WriteString("\u001b[42m" + indent + line + "\u001b[K\n")
				}

				console.Stdout.WriteString("\u001b[0m")

			} else {

				for _, line := range lines {
					console.Stdout.WriteString(indent + line + "\n")
				}

			}

		}

		console.mutex.Lock()
		console.Messages = append(console.Messages, message)
		console.mutex.Unlock()

	}

}

func (console *Console) Progress(raw string) {

	// Multiline is forbidden
	if strings.Contains(raw, "\n") {
		raw = raw[0:strings.Index(raw, "\n")]
	}

	message := NewConsoleMessage("Progress", raw)

	if enabled, _ := console.methods["Progress"]; enabled == true {

		indent := toConsoleIndent(console.offset)
		lines  := message.Lines()

		if len(lines) == 1 {

			console.mutex.Lock()

			if len(console.Messages) > 0 {

				last_progress := -1

				for m := len(console.Messages) - 1; m >= 0; m-- {

					if console.Messages[m].Method == "Progress" {
						last_progress = m
						break
					}

				}

				if last_progress != -1 && last_progress == len(console.Messages) - 1 {

					last_message := console.Messages[last_progress]

					if message.IsSame(last_message) {

						if console.Stdout != nil {

							if console.use_colors == true {
								console.Stdout.WriteString("\033[A\033[2K\r")
								console.Stdout.WriteString("\u001b[40m" + indent + lines[0] + "\u001b[K\u001b[0m\n")
							} else {
								console.Stdout.WriteString("\033[A\033[2K\r")
								console.Stdout.WriteString(indent + lines[0] + "\n")
							}

						}

						console.Messages[len(console.Messages) - 1] = message

					} else {

						if console.Stdout != nil {

							if console.use_colors == true {
								console.Stdout.WriteString("\u001b[40m" + indent + lines[0] + "\u001b[K\u001b[0m\n")
							} else {
								console.Stdout.WriteString(indent + lines[0] + "\n")
							}

						}

						console.Messages = append(console.Messages, message)

					}

				} else {

					if console.Stdout != nil {

						if console.use_colors == true {
							console.Stdout.WriteString("\u001b[40m" + indent + lines[0] + "\u001b[K\u001b[0m\n")
						} else {
							console.Stdout.WriteString(indent + lines[0] + "\n")
						}

					}

					console.Messages = append(console.Messages, message)

				}

			} else {

				if console.Stdout != nil {

					if console.use_colors == true {
						console.Stdout.WriteString("\u001b[40m" + indent + lines[0] + "\u001b[K\u001b[0m\n")
					} else {
						console.Stdout.WriteString(indent + lines[0] + "\n")
					}

				}

				console.Messages = append(console.Messages, message)

			}

			console.mutex.Unlock()

		}

	}

}

func (console *Console) Warn(raw string) {

	message := NewConsoleMessage("Warn", raw)

	if enabled, _ := console.methods["Warn"]; enabled == true {

		indent := toConsoleIndent(console.offset)
		lines  := message.Lines()

		if len(lines) > 0 && console.Stdout != nil {

			if console.use_colors == true {

				for _, line := range lines {
					console.Stdout.WriteString("\u001b[43m" + indent + line + "\u001b[K\n")
				}

				console.Stdout.WriteString("\u001b[0m")

			} else {

				for _, line := range lines {
					console.Stdout.WriteString(indent + line + "\n")
				}

			}

		}

		console.mutex.Lock()
		console.Messages = append(console.Messages, message)
		console.mutex.Unlock()

	}

}

