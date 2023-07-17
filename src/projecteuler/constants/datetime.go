package constants

const SUNDAY = 0
const MONDAY = 1
const TUESDAY = 2
const WEDNESDAY = 3
const THURSDAY = 4
const FRIDAY = 5
const SATURDAY = 6

const JANUARY = 0
const FEBRUARY = 1
const MARCH = 2
const APRIL = 3
const MAY = 4
const JUNE = 5
const JULY = 6
const AUGUST = 7
const SEPTEMBER = 8
const OCTOBER = 9
const NOVEMBER = 10
const DECEMBER = 11

var MONTHS_OF_YEAR = []string{
	"JANUARY",
	"FEBRUARY",
	"MARCH",
	"APRIL",
	"MAY",
	"JUNE",
	"JULY",
	"AUGUST",
	"SEPTEMBER",
	"OCTOBER",
	"NOVEMBER",
	"DECEMBER",
}

// Golang maps are *UNSORTED* (randomly sorted in runtime),
// so it is better to iterate MONTHS_OF_YEAR as keys of this map
var DAYS_IN_MONTH = map[string]int{
	"JANUARY":   31,
	"FEBRUARY":  28,
	"MARCH":     31,
	"APRIL":     30,
	"MAY":       31,
	"JUNE":      30,
	"JULY":      31,
	"AUGUST":    31,
	"SEPTEMBER": 30,
	"OCTOBER":   31,
	"NOVEMBER":  30,
	"DECEMBER":  31,
}
