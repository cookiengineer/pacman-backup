package types

import "bytes"
import "strconv"
import "strings"
import gotime "time"

type Time struct {
	Hour   uint `json:"hour"`
	Minute uint `json:"minute"`
	Second uint `json:"second"`
	valid  bool
}

func NewTime(value string) Time {

	var time Time

	if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":")

		if len(chunks) == 3 {

			time.Parse(value)

			if value == "00:00:00" && time.Hour == 0 && time.Minute == 0 && time.Second == 0 {
				time.valid = true
			} else if time.Hour != 0 || time.Minute != 0 || time.Second != 0 {
				time.valid = true
			}

		} else if len(chunks) == 2 {

			time.Parse(value)

			if value == "00:00" && time.Hour == 0 && time.Minute == 0 && time.Second == 0 {
				time.valid = true
			} else if time.Hour != 0 || time.Minute != 0 {
				time.valid = true
			}

		}

	}

	return time

}

func (time *Time) AddHour() {

	time.Hour = time.Hour + 1

}

func (time *Time) AddMinute() {

	var hour uint = time.Hour
	var minute uint = time.Minute + 1

	if minute > 60 {
		hour += 1
		minute -= 60
	}

	time.Hour = hour
	time.Minute = minute

}

func (time *Time) AddSecond() {

	var hour uint = time.Hour
	var minute uint = time.Minute
	var second uint = time.Second + 1

	if second > 60 {
		minute += 1
		second -= 60
	}

	if minute > 60 {
		hour += 1
		minute -= 60
	}

	time.Hour = hour
	time.Minute = minute
	time.Second = second

}

func (time *Time) AddTime(other Time) {

	if time.IsBefore(other) {

		var hour uint = time.Hour + other.Hour
		var minute uint = time.Minute + other.Minute
		var second uint = time.Second + other.Second

		if second > 60 {
			minute += 1
			second -= 60
		}

		if minute > 60 {
			hour += 1
			minute -= 60
		}

		time.Hour = hour
		time.Minute = minute
		time.Second = second

	}

}

func (time *Time) IsAfter(other Time) bool {

	var result bool = false

	if time.Hour > other.Hour {
		result = true
	} else if time.Hour == other.Hour {

		if time.Minute > other.Minute {
			result = true
		} else if time.Minute == other.Minute {

			if time.Second > other.Second {
				result = true
			} else if time.Second == other.Second {
				result = false
			}

		}

	}

	return result

}

func (time *Time) IsBefore(other Time) bool {

	var result bool = false

	if time.Hour < other.Hour {
		result = true
	} else if time.Hour == other.Hour {

		if time.Minute < other.Minute {
			result = true
		} else if time.Minute == other.Minute {

			if time.Second < other.Second {
				result = true
			} else if time.Second == other.Second {
				result = false
			}

		}

	}

	return result

}

func (time *Time) IsPast() bool {

	now := NewTime(gotime.Now().Format(gotime.TimeOnly))

	return time.IsBefore(now)

}

func (time *Time) IsFuture() bool {

	now := NewTime(gotime.Now().Format(gotime.TimeOnly))

	return time.IsAfter(now)

}

func (time *Time) IsValid() bool {
	return time.valid
}

func (time *Time) Offset(offset string) {

	var operator string
	var offset_hours uint64
	var offset_minutes uint64
	var err_hours error
	var err_minutes error

	if len(offset) == 6 {

		if strings.HasPrefix(offset, "+") || strings.HasPrefix(offset, "-") {
			operator = string(offset[0])
			offset_hours, err_hours = strconv.ParseUint(offset[1:3], 10, 64)
			offset_minutes, err_minutes = strconv.ParseUint(offset[4:], 10, 64)
		}

	} else if len(offset) == 5 {

		if strings.HasPrefix(offset, "+") || strings.HasPrefix(offset, "-") {
			operator = string(offset[0])
			offset_hours, err_hours = strconv.ParseUint(offset[1:3], 10, 64)
			offset_minutes, err_minutes = strconv.ParseUint(offset[3:], 10, 64)
		}

	}

	if err_hours == nil && err_minutes == nil {

		if operator == "+" {

			var hour = int(time.Hour - uint(offset_hours))
			var minute = int(time.Minute - uint(offset_minutes))

			if minute < 0 {
				hour -= 1
				minute += 60
			}

			if hour < 0 {
				hour += 24
			}

			time.Hour = uint(hour)
			time.Minute = uint(minute)

		} else if operator == "-" {

			var hour = int(time.Hour + uint(offset_hours))
			var minute = int(time.Minute + uint(offset_minutes))

			if minute > 60 {
				hour += 1
				minute -= 60
			}

			if hour > 24 {
				hour -= 24
			}

			time.Hour = uint(hour)
			time.Minute = uint(minute)

		}

	}

}

func (time *Time) Parse(value string) {

	if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":")

		if len(chunks) == 3 {

			num1, err1 := strconv.ParseUint(chunks[0], 10, 64)
			num2, err2 := strconv.ParseUint(chunks[1], 10, 64)
			num3, err3 := strconv.ParseUint(chunks[2], 10, 64)

			if err1 == nil {

				if num1 >= 0 && num1 <= 24 {
					time.Hour = uint(num1)
				}

			}

			if err2 == nil {

				if num2 >= 0 && num2 <= 60 {
					time.Minute = uint(num2)
				}

			}

			if err3 == nil {

				if num3 >= 0 && num3 <= 60 {
					time.Second = uint(num3)
				}

			}

		} else if len(chunks) == 2 {

			num1, err1 := strconv.ParseUint(chunks[0], 10, 64)
			num2, err2 := strconv.ParseUint(chunks[1], 10, 64)

			if err1 == nil {

				if num1 >= 0 && num1 <= 24 {
					time.Hour = uint(num1)
				}

			}

			if err2 == nil {

				if num2 >= 0 && num2 <= 60 {
					time.Minute = uint(num2)
				}

			}

		}

	}

}

func (time *Time) String() string {

	var buffer bytes.Buffer

	if time.Hour > 99 {
		buffer.WriteString(formatUint(time.Hour, 0))
	} else {
		buffer.WriteString(formatUint(time.Hour, 2))
	}

	buffer.WriteString(":")
	buffer.WriteString(formatUint(time.Minute, 2))
	buffer.WriteString(":")
	buffer.WriteString(formatUint(time.Second, 2))

	return buffer.String()

}

func (time *Time) ToZulu() {

	var offset = strings.Split(gotime.Now().String(), " ")[2]

	if len(offset) == 5 {
		time.Offset(offset)
	}

}

func (time Time) MarshalJSON() ([]byte, error) {

	quoted := strconv.Quote(time.String())

	return []byte(quoted), nil

}

func (time *Time) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	check := NewTime(unquoted)

	if check.IsValid() {

		time.Hour = check.Hour
		time.Minute = check.Minute
		time.Second = check.Second

	}

	return nil

}
