package types

import "bytes"
import "strconv"
import "strings"
import "time"

var WEEKDAYS = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
var MONTHS = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var MONTHDAYS = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var TIMEZONES = []string{

	"BST", "BDT", // Bering
	"HST", "HDT", "HWT", "HPT", // Hawaii
	"AHST", "AHDT", "AKST", "AKDT", // Alaska
	"NST", "NDT", "NWT", "NPT", // Nome
	"YST", "YDT", "YWT", "YPT", "YDDT", // Yukon
	"PST", "PDT", "PWT", "PPT", "PDDT", // Pacific
	"MST", "MDT", "MWT", "MPT", "MDDT", // Mountain
	"CST", "CDT", "CWT", "CPT", "CDDT", // Central America
	"EST", "EDT", "EWT", "EPT", "EDDT", // Eastern America
	"NST", "NDT", "NWT", "NPT", "NDDT", // Newfoundland
	"AST", "ADT", "APT", "AWT", "ADDT", // Atlantic

	"GMT", "BST", "IST", "BDST", // Great Britain
	"WET", "WEST", "WEMT", // Western Europe
	"CET", "CEST", "CEMT", // Central Europe
	"MET", "MEST", // Middle Europe
	"EET", "EEST", // Eastern Europe

	"WAT", "WAST", // Western Africa
	"CAT", "CAST", // Central Africa
	"EAT",  // Eastern Africa
	"SAST", // South Africa

	"MSK", "MSD", // Moscow
	"IST", "IDT", "IDDT", // Israel
	"CST", "CDT", // China
	"PKT", "PKST", // Pakistan
	"IST",                // India
	"HKT", "HKST", "HPT", // Hong Kong
	"KST", "KDT", // Korea
	"JST", "JDT", // Japan

	"AWST", "AWDT", // Western Australia
	"ACST", "ACDT", // Central Australia
	"AEST", "AEDT", // Eastern Australia
	"WIB", "WIT", "WITA", // Waktu Indonesia Barat/Timur/Tengah
	"PST", "PDT", // Philippines
	"GST", "GDT", "CHST", // Guam / Chamorro
	"NZST", "NZDT", // New Zealand
	"SST", // Samoa

	"UTC", // Universal

}

func isDay(value string) bool {

	var result bool = false

	num, err := strconv.ParseUint(value, 10, 64)

	if err == nil {

		if num >= 1 && num <= 31 {
			result = true
		}

	}

	return result

}

func isISO8601(value string) bool {

	var result bool = false

	// "2006-05-30T10:02:00"
	// "2006-05-30T10:02:00.000"
	// "2006-05-30T10:02:00.000Z"
	// "2006-05-30 10:02:00"

	if strings.Contains(value, "T") && strings.HasSuffix(value, "Z") {

		var date = strings.Split(value, "T")[0]
		var time = strings.Split(value, "T")[1]

		if strings.HasSuffix(time, "Z") {
			time = time[0 : len(time)-1]
		}

		if strings.Contains(time, "+") {
			time = strings.Split(time, "+")[0]
		} else if strings.Contains(time, "-") {
			time = strings.Split(time, "-")[0]
		}

		var check_date = strings.Split(date, "-")
		var check_time = strings.Split(time, ":")

		if len(check_date) == 3 && len(check_time) == 3 {
			result = true
		}

	} else if strings.Contains(value, "T") {

		var date = strings.Split(value, "T")[0]
		var time = strings.Split(value, "T")[1]

		if strings.Contains(time, "+") {
			time = strings.Split(time, "+")[0]
		} else if strings.Contains(time, "-") {
			time = strings.Split(time, "-")[0]
		}

		var check_date = strings.Split(date, "-")
		var check_time = strings.Split(time, ":")

		if len(check_date) == 3 && len(check_time) == 3 {
			result = true
		}

	}

	return result

}

func isISO8601Date(value string) bool {

	var result bool = false

	if strings.Contains(value, "-") {

		var chunks = strings.Split(value, "-")

		if len(chunks) == 3 {
			result = true
		}

	}

	return result

}

func isYYYYMMDD(value string) bool {

	var result bool = false

	// 20060530

	if len(value) == 8 && strings.Contains(value, "-") == false {

		num1, err1 := strconv.ParseUint(value[0:4], 10, 64)
		num2, err2 := strconv.ParseUint(value[4:6], 10, 64)
		num3, err3 := strconv.ParseUint(value[6:8], 10, 64)

		if err1 == nil && num1 >= 1752 {

			if err2 == nil && num2 >= 1 && num2 <= 12 {

				if err3 == nil && num3 >= 1 && num3 <= 31 {
					result = true
				}

			}

		}

	}

	return result

}

func isYYYYMM(value string) bool {

	var result bool = false

	// 2006-05
	// 200605

	if len(value) == 7 && strings.Contains(value, "-") {

		year := strings.Split(value, "-")[0]
		month := strings.Split(value, "-")[1]

		num1, err1 := strconv.ParseUint(year, 10, 64)
		num2, err2 := strconv.ParseUint(month, 10, 64)

		if err1 == nil && num1 >= 1752 {

			if err2 == nil && num2 >= 1 && num2 <= 12 {
				result = true
			}

		}

	} else if len(value) == 6 && strings.Contains(value, "-") == false {

		num1, err1 := strconv.ParseUint(value[0:4], 10, 64)
		num2, err2 := strconv.ParseUint(value[4:6], 10, 64)

		if err1 == nil && num1 >= 1752 {

			if err2 == nil && num2 >= 1 && num2 <= 12 {
				result = true
			}

		}

	}

	return result

}

func isMeridiem(value string) bool {

	if value == "AM" || value == "PM" {
		return true
	}

	return false

}

func isMonth(value string) bool {

	var result bool = false

	for m := 0; m < len(MONTHS); m++ {

		if MONTHS[m] == value {
			result = true
			break
		}

	}

	return result

}

func isTime(value string) bool {

	var result bool = false

	if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":")

		if len(chunks) == 3 {
			result = true
		} else if len(chunks) == 2 {
			result = true
		}

	}

	return result

}

func isTimezone(value string) bool {

	var result bool = false

	for t := 0; t < len(TIMEZONES); t++ {

		if TIMEZONES[t] == value {
			result = true
			break
		}

	}

	return result

}

func isWeekday(value string) bool {

	var result bool = false

	for w := 0; w < len(WEEKDAYS); w++ {

		if WEEKDAYS[w] == value {
			result = true
			break
		}

	}

	return result

}

func isLeapYear(value uint) bool {

	if value%4 != 0 {
		return false
	} else if value%100 != 0 {
		return true
	} else if value%400 != 0 {
		return false
	} else {
		return true
	}

}

func isYear(value string) bool {

	var result bool = false

	num, err := strconv.ParseUint(value, 10, 64)

	if err == nil {

		if num >= 1752 {
			result = true
		}

	}

	return result

}

func parseDay(datetime *Datetime, value string) {

	num, err := strconv.ParseUint(value, 10, 64)

	if err == nil {

		if num >= 1 && num <= 31 {
			datetime.Day = uint(num)
		}

	}

}

func parseISO8601(datetime *Datetime, value string) {

	if strings.HasSuffix(value, "Z") {
		value = value[0 : len(value)-1]
	}

	if strings.Contains(value, "T") {

		var date = strings.Split(value, "T")[0]
		var time = strings.Split(value, "T")[1]
		var tmp = strings.Split(date, "-")

		if len(tmp) == 3 {
			parseYear(datetime, tmp[0])
			parseMonth(datetime, tmp[1])
			parseDay(datetime, tmp[2])
		}

		// Strip out milliseconds
		if strings.Contains(time, ".") {
			time = strings.Split(time, ".")[0]
		}

		if strings.Contains(time, ":") {
			parseTime(datetime, time)
		}

	}

}

func parseISO8601Date(datetime *Datetime, value string) {

	if strings.Contains(value, "-") {

		var tmp = strings.Split(value, "-")

		if len(tmp) == 3 {
			parseYear(datetime, tmp[0])
			parseMonth(datetime, tmp[1])
			parseDay(datetime, tmp[2])
		}

	}

}

func parseYYYYMM(datetime *Datetime, value string) {

	if len(value) == 7 && strings.Contains(value, "-") {

		year := strings.Split(value, "-")[0]
		month := strings.Split(value, "-")[1]

		parseYear(datetime, year)
		parseMonth(datetime, month)
		datetime.Day = 1

	} else if len(value) == 6 && strings.Contains(value, "-") == false {

		parseYear(datetime, value[0:4])
		parseMonth(datetime, value[4:6])
		datetime.Day = 1

	}

}

func parseYYYYMMDD(datetime *Datetime, value string) {

	if len(value) == 8 {
		parseYear(datetime, value[0:4])
		parseMonth(datetime, value[4:6])
		parseDay(datetime, value[6:8])
	}

}

func parseMonth(datetime *Datetime, value string) {

	for m := 0; m < len(MONTHS); m++ {

		if MONTHS[m] == value {
			datetime.Month = uint(m + 1)
			break
		}

	}

	if datetime.Month == 0 {

		num, err := strconv.ParseUint(value, 10, 64)

		if err == nil {

			if num >= 1 && num <= 12 {
				datetime.Month = uint(num)
			}

		}

	}

}

func parseTime(datetime *Datetime, value string) {

	if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":")

		if len(chunks) == 3 {

			num1, err1 := strconv.ParseUint(chunks[0], 10, 64)
			num2, err2 := strconv.ParseUint(chunks[1], 10, 64)
			num3, err3 := strconv.ParseUint(chunks[2], 10, 64)

			if err1 == nil {

				if num1 >= 0 && num1 <= 24 {
					datetime.Hour = uint(num1)
				}

			}

			if err2 == nil {

				if num2 >= 0 && num2 <= 60 {
					datetime.Minute = uint(num2)
				}

			}

			if err3 == nil {

				if num3 >= 0 && num3 <= 60 {
					datetime.Second = uint(num3)
				}

			}

		} else if len(chunks) == 2 {

			num1, err1 := strconv.ParseUint(chunks[0], 10, 64)
			num2, err2 := strconv.ParseUint(chunks[1], 10, 64)

			if err1 == nil {

				if num1 >= 0 && num1 <= 24 {
					datetime.Hour = uint(num1)
				}

			}

			if err2 == nil {

				if num2 >= 0 && num2 <= 60 {
					datetime.Minute = uint(num2)
				}

			}

		}

	}

}

func parseYear(datetime *Datetime, value string) {

	num, err := strconv.ParseUint(value, 10, 64)

	if err == nil {

		if num >= 1970 {
			datetime.Year = uint(num)
		}

	}

}

func toChunks(value string) []string {

	var chunks []string
	var values []string = strings.Split(strings.TrimSpace(value), " ")

	for v := 0; v < len(values); v++ {

		var value = strings.TrimSpace(values[v])

		if value != "" {
			chunks = append(chunks, value)
		}

	}

	return chunks

}

type Datetime struct {
	Year   uint // `json:"year"`
	Month  uint // `json:"month"`
	Day    uint // `json:"day"`
	Hour   uint // `json:"hour"`
	Minute uint // `json:"minute"`
	Second uint // `json:"second"`
}

func NewDatetime() Datetime {

	var datetime Datetime

	datetime.Parse(time.Now().Format(time.RFC3339))

	return datetime

}

func ToDatetime(value string) Datetime {

	var datetime Datetime

	datetime.Parse(value)

	return datetime

}

func (datetime *Datetime) IsAfter(other Datetime) bool {

	var result bool = false

	if datetime.Year > other.Year {
		result = true
	} else if datetime.Year == other.Year {

		if datetime.Month > other.Month {
			result = true
		} else if datetime.Month == other.Month {

			if datetime.Day > other.Day {
				result = true
			} else if datetime.Day == other.Day {

				if datetime.Hour > other.Hour {
					result = true
				} else if datetime.Hour == other.Hour {

					if datetime.Minute > other.Minute {
						result = true
					} else if datetime.Minute == other.Minute {

						if datetime.Second > other.Second {
							result = true
						} else if datetime.Second == other.Second {
							result = false
						}

					}

				}

			}

		}

	}

	return result

}

func (datetime *Datetime) IsBefore(other Datetime) bool {

	var result bool = false

	if datetime.Year < other.Year {
		result = true
	} else if datetime.Year == other.Year {

		if datetime.Month < other.Month {
			result = true
		} else if datetime.Month == other.Month {

			if datetime.Day < other.Day {
				result = true
			} else if datetime.Day == other.Day {

				if datetime.Hour < other.Hour {
					result = true
				} else if datetime.Hour == other.Hour {

					if datetime.Minute < other.Minute {
						result = true
					} else if datetime.Minute == other.Minute {

						if datetime.Second < other.Second {
							result = true
						} else if datetime.Second == other.Second {
							result = false
						}

					}

				}

			}

		}

	}

	return result

}

func (datetime *Datetime) IsFuture() bool {

	now := ToDatetime(time.Now().Format(time.RFC3339))

	return datetime.IsAfter(now)

}

func (datetime *Datetime) IsPast() bool {

	now := ToDatetime(time.Now().Format(time.RFC3339))

	return datetime.IsBefore(now)

}

func (datetime *Datetime) IsValid() bool {

	if datetime.Year > 1752 {

		if datetime.Month >= 1 && datetime.Month <= 12 {

			if datetime.Day >= 1 && datetime.Day <= 31 {
				return true
			}

		}

	}

	return false

}

func (datetime *Datetime) Offset(offset string) {

	var operator string
	var offset_hours uint64
	var offset_minutes uint64
	var err_hours error
	var err_minutes error

	if len(offset) == 6 && datetime.Year > 0 && datetime.Month > 0 && datetime.Day > 0 {

		if strings.HasPrefix(offset, "+") || strings.HasPrefix(offset, "-") {
			operator = string(offset[0])
			offset_hours, err_hours = strconv.ParseUint(offset[1:3], 10, 64)
			offset_minutes, err_minutes = strconv.ParseUint(offset[4:], 10, 64)
		}

	} else if len(offset) == 5 && datetime.Year > 0 && datetime.Month > 0 && datetime.Day > 0 {

		if strings.HasPrefix(offset, "+") || strings.HasPrefix(offset, "-") {
			operator = string(offset[0])
			offset_hours, err_hours = strconv.ParseUint(offset[1:3], 10, 64)
			offset_minutes, err_minutes = strconv.ParseUint(offset[3:], 10, 64)
		}

	}

	if err_hours == nil && err_minutes == nil {

		if operator == "+" {

			var year = datetime.Year
			var month = datetime.Month
			var day = datetime.Day
			var hour = int(datetime.Hour - uint(offset_hours))
			var minute = int(datetime.Minute - uint(offset_minutes))

			if minute < 0 {
				hour -= 1
				minute += 60
			}

			if hour < 0 {
				day -= 1
				hour += 24
			}

			if day <= 0 {

				if month > 1 {

					month -= 1

					if isLeapYear(year) && month == 2 {
						day += uint(MONTHDAYS[month-1]) + 1
					} else {
						day += uint(MONTHDAYS[month-1])
					}

				} else {

					year -= 1
					month = 12
					day = 31

				}

			}

			datetime.Year = year
			datetime.Month = month
			datetime.Day = day
			datetime.Hour = uint(hour)
			datetime.Minute = uint(minute)

		} else if operator == "-" {

			var year = datetime.Year
			var month = datetime.Month
			var day = datetime.Day
			var hour = datetime.Hour + uint(offset_hours)
			var minute = datetime.Minute + uint(offset_minutes)

			if minute > 60 {
				hour += 1
				minute -= 60
			}

			if hour > 24 {
				day += 1
				hour -= 24
			}

			var monthdays = uint(MONTHDAYS[month-1])

			if isLeapYear(year) && month == 2 {
				monthdays += 1
			}

			if day > monthdays {
				month += 1
				day -= monthdays
			}

			if month > 12 {
				year += 1
				month -= 12
			}

			datetime.Year = year
			datetime.Month = month
			datetime.Day = day
			datetime.Hour = hour
			datetime.Minute = minute

		}

	}

}

func (datetime *Datetime) Yesterday() Datetime {

	var result Datetime

	if datetime.Day == 1 {

		if datetime.Month == 1 {
			result.Year = datetime.Year - 1
			result.Month = 12
			result.Day = result.ToDays()
		} else {
			result.Year = datetime.Year
			result.Month = datetime.Month - 1
			result.Day = result.ToDays()
		}

	} else {
		result.Year = datetime.Year
		result.Month = datetime.Month
		result.Day = datetime.Day - 1
	}

	return result

}

func (datetime *Datetime) Tomorrow() Datetime {

	var result Datetime

	if datetime.Day == datetime.ToDays() {

		if datetime.Month == 12 {
			result.Year = datetime.Year + 1
			result.Month = 1
			result.Day = 1
		} else {
			result.Year = datetime.Year
			result.Month = datetime.Month + 1
			result.Day = 1
		}

	} else {
		result.Year = datetime.Year
		result.Month = datetime.Month
		result.Day = datetime.Day + 1
	}

	return result

}

func (datetime *Datetime) Parse(value string) {

	var chunks []string = toChunks(value)
	var isZulu bool = false

	if isISO8601(strings.TrimSpace(value)) {

		if strings.Contains(chunks[0], "T") {

			var time_suffix = strings.Split(chunks[0], "T")[1]

			if strings.HasSuffix(time_suffix, "Z") {

				parseISO8601(datetime, chunks[0])
				isZulu = true

			} else if strings.Contains(time_suffix, "+") {

				parseISO8601(datetime, chunks[0][0:19])

				datetime.Offset(strings.Split(time_suffix, "+")[1])
				isZulu = true

			} else if strings.Contains(time_suffix, "-") {

				parseISO8601(datetime, chunks[0][0:19])
				datetime.Offset(strings.Split(time_suffix, "-")[1])
				isZulu = true

			}

		}

	} else if len(chunks) == 1 {

		if isISO8601Date(chunks[0]) {

			parseISO8601Date(datetime, chunks[0])
			isZulu = true

		} else if isYYYYMMDD(chunks[0]) {

			parseYYYYMMDD(datetime, chunks[0])
			isZulu = true

		} else if isYYYYMM(chunks[0]) {

			parseYYYYMM(datetime, chunks[0])
			isZulu = true

		}

	} else if len(chunks) == 2 {

		if isISO8601Date(chunks[0]) && isTime(chunks[1]) {

			parseISO8601Date(datetime, chunks[0])
			parseTime(datetime, chunks[1])
			isZulu = true

		}

	} else if len(chunks) == 3 {

		if isMonth(chunks[0]) && isDay(chunks[1]) && isTime(chunks[2]) {

			datetime.Year = uint(time.Now().Year())

			parseMonth(datetime, chunks[0])
			parseDay(datetime, chunks[1])
			parseTime(datetime, chunks[2])

		}

	} else if len(chunks) == 5 {

		if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isYear(chunks[4]) {

			parseMonth(datetime, chunks[1])
			parseDay(datetime, chunks[2])
			parseTime(datetime, chunks[3])
			parseYear(datetime, chunks[4])

		}

	} else if len(chunks) == 6 {

		if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isMeridiem(chunks[4]) && isYear(chunks[5]) {

			parseMonth(datetime, chunks[1])
			parseDay(datetime, chunks[2])
			parseTime(datetime, chunks[3])
			parseYear(datetime, chunks[5])

			if chunks[4] == "PM" {
				datetime.Hour += 12
			}

		} else if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isTimezone(chunks[4]) && isYear(chunks[5]) {

			parseMonth(datetime, chunks[1])
			parseDay(datetime, chunks[2])
			parseTime(datetime, chunks[3])
			parseYear(datetime, chunks[5])

		}

	} else if len(chunks) == 7 {

		if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isMeridiem(chunks[4]) && isTimezone(chunks[5]) && isYear(chunks[6]) {

			parseMonth(datetime, chunks[1])
			parseDay(datetime, chunks[2])
			parseTime(datetime, chunks[3])
			parseYear(datetime, chunks[6])

			if chunks[4] == "PM" {
				datetime.Hour += 12
			}

		}

	}

	if isZulu == false {
		datetime.ToZulu()
	}

}

func (datetime *Datetime) String() string {

	var buffer bytes.Buffer

	buffer.WriteString(formatUint(datetime.Year, 4))
	buffer.WriteString("-")
	buffer.WriteString(formatUint(datetime.Month, 2))
	buffer.WriteString("-")
	buffer.WriteString(formatUint(datetime.Day, 2))
	buffer.WriteString(" ")
	buffer.WriteString(formatUint(datetime.Hour, 2))
	buffer.WriteString(":")
	buffer.WriteString(formatUint(datetime.Minute, 2))
	buffer.WriteString(":")
	buffer.WriteString(formatUint(datetime.Second, 2))

	return buffer.String()

}

func (datetime *Datetime) ToDays() uint {

	var days uint

	var month = datetime.Month
	var year = datetime.Year

	if isLeapYear(year) && month == 2 {
		days = uint(MONTHDAYS[month-1]) + 1
	} else {
		days = uint(MONTHDAYS[month-1])
	}

	return days

}

func (datetime *Datetime) ToDatetimeDifference(other Datetime) Datetime {

	var result Datetime

	if datetime.IsBefore(other) {

		var years uint = 0
		var months uint = 0
		var days uint = 0
		var hours uint = (24 - datetime.Hour) + other.Hour
		var minutes uint = (60 - datetime.Minute) + other.Minute
		var seconds uint = (60 - datetime.Second) + other.Second

		tmp := ToDatetime(datetime.String())

		if tmp.Year <= other.Year {

			if tmp.Month <= other.Month {

				for tmp.Year < other.Year {
					years += 1
					tmp.Year += 1
				}

				for tmp.Month < other.Month {
					months += 1
					tmp.Month += 1
				}

				if tmp.Day <= other.Day {
					days = other.Day - tmp.Day
				} else {
					months -= 1
					days = (tmp.ToDays() - tmp.Day) + other.Day
				}

			} else {

				if tmp.Day <= other.Day {
					months = (12 - tmp.Month) + other.Month
					days = other.Day - tmp.Day
				} else {
					months = (12 - tmp.Month) + other.Month - 1
					days = (tmp.ToDays() - tmp.Day) + other.Day
				}

			}

		}

		if hours > 0 || minutes > 0 || seconds > 0 {
			days -= 1
		}

		if minutes > 0 || seconds > 0 {

			hours -= 1

			if seconds > 60 {
				seconds -= 60
				minutes += 1
			}

			if seconds > 60 {
				seconds -= 60
				minutes += 1
			}

			if minutes > 60 {
				minutes -= 60
				hours += 1
			}

			if minutes > 60 {
				minutes -= 60
				hours += 1
			}

			if hours == 23 && minutes == 60 && seconds == 60 {
				days += 1
				hours = 0
				minutes = 0
				seconds = 0
			} else if minutes == 60 && seconds == 60 {
				seconds = 0
				minutes = 0
				hours += 1
			} else if minutes == 60 {
				minutes = 0
				hours += 1
			}

		}

		result.Year = years
		result.Month = months
		result.Day = days
		result.Hour = hours
		result.Minute = minutes
		result.Second = seconds

	}

	return result

}

func (datetime *Datetime) ToTimeDifference(other Datetime) Time {

	var result Time

	if datetime.IsBefore(other) {

		var days uint = 0
		var hours uint = (24 - datetime.Hour) + other.Hour
		var minutes uint = (60 - datetime.Minute) + other.Minute
		var seconds uint = (60 - datetime.Second) + other.Second

		tmp := ToDatetime(datetime.String())

		if tmp.Year != other.Year {

			for tmp.Year < other.Year-1 {

				if isLeapYear(tmp.Year) {
					days += 366
				} else {
					days += 365
				}

				tmp.Year += 1

			}

			days += (tmp.ToDays() - tmp.Day)
			tmp.Day = tmp.ToDays()
			tmp.Month += 1

			for tmp.Month <= 12 {
				days += tmp.ToDays()
				tmp.Month += 1
			}

			days += 1
			tmp.Year += 1
			tmp.Month = 1
			tmp.Day = 1

		}

		if tmp.Year == other.Year {

			if tmp.Month == other.Month {

				if tmp.Day != other.Day {
					days += other.Day - tmp.Day
				}

			} else if tmp.Month < other.Month {

				days += (tmp.ToDays() - tmp.Day)
				tmp.Month += 1

				for tmp.Month != other.Month {
					days += tmp.ToDays()
					tmp.Month += 1
				}

				days += other.Day

			}

		}

		if days > 0 {
			hours += (days - 1) * 24
		} else {
			hours -= 24
		}

		if minutes > 0 || seconds > 0 {

			hours -= 1

			if seconds > 60 {
				seconds -= 60
				minutes += 1
			}

			if seconds > 60 {
				seconds -= 60
				minutes += 1
			}

			if minutes > 60 {
				minutes -= 60
				hours += 1
			}

			if minutes > 60 {
				minutes -= 60
				hours += 1
			}

			if minutes == 60 && seconds == 60 {
				seconds = 0
				minutes = 0
				hours += 1
			} else if minutes == 60 {
				minutes = 0
				hours += 1
			}

		}

		result.Hour = hours
		result.Minute = minutes
		result.Second = seconds

	}

	return result

}

func (datetime *Datetime) ToWeekday() string {

	// month >= 1 && month <= 12
	// year > 1752

	// Don't ask. Just Don't.

	var year = datetime.Year
	var month = datetime.Month
	var day = datetime.Day
	var offsets []uint = []uint{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}

	if month < 3 {
		year -= 1
	}

	var index uint = (year + year/4 - year/100 + year/400 + offsets[month-1] + day) % 7
	var weekdays []string = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	return weekdays[index]

}

func (datetime *Datetime) ToZulu() {

	var offset = strings.Split(time.Now().String(), " ")[2]

	if len(offset) == 5 && datetime.Year > 0 && datetime.Month > 0 && datetime.Day > 0 {
		datetime.Offset(offset)
	}

}

func (datetime Datetime) MarshalJSON() ([]byte, error) {

	quoted := strconv.Quote(datetime.String())

	return []byte(quoted), nil

}

func (datetime *Datetime) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	check := ToDatetime(unquoted)

	if check.IsValid() {

		datetime.Year = check.Year
		datetime.Month = check.Month
		datetime.Day = check.Day
		datetime.Hour = check.Hour
		datetime.Minute = check.Minute
		datetime.Second = check.Second

	}

	return nil

}
