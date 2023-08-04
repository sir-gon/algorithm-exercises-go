/**
 *
 * https://projecteuler.net/problem=XX
 *
 *
 */

package projecteuler

import (
	constants "gon.cl/algorithm-exercises/src/projecteuler/constants"
	utils "gon.cl/algorithm-exercises/src/projecteuler/utils"
)

func Problem0019(
	_dayOfWeek int,
	_sinceYear int,
	_untilYear int) int {

	const initYear = 1900
	var resultCount = 0
	var accumulated_remainder = 0
	var excess = 0

	const __FEBRUARY__KEY__ = "FEBRUARY"
	// daysInMonth := constants.DAYS_IN_MONTH

	for y := initYear; y <= _untilYear; y++ {
		var leap int

		if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
			utils.Debug("Year %d has leap-day", y)
			leap = 1
		} else {
			leap = 0
		}
		constants.DAYS_IN_MONTH[__FEBRUARY__KEY__] = 28 + leap

		for _, month := range constants.MONTHS_OF_YEAR {
			days := constants.DAYS_IN_MONTH[month]
			utils.Debug("Year %d| Month: %s | days %d", y, month, days)

			accumulated_remainder += days % 7
			if accumulated_remainder%7 == _dayOfWeek {
				if y <= _sinceYear {
					excess += 1
				}
				resultCount += 1
			}
		}
	}

	utils.Info("Problem0019 result: (%d - %d) => %d", resultCount, excess, resultCount-excess)

	return resultCount - excess
}
