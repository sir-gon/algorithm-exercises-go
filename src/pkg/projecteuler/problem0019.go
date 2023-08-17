/**
 * @link Problem definition [[docs/projecteuler/problem0019.md]]
 */

package projecteuler

import (
	constants "gon.cl/algorithm-exercises/src/pkg/projecteuler/constants"
	utils "gon.cl/algorithm-exercises/src/utils"
)

const __FEBRUARY_KEY__ = "FEBRUARY"

func Problem0019(
	dayOfWeek int,
	sinceYear int,
	untilYear int) int {

	const initYear = 1900
	var resultCount = 0
	var accumulatedRemainder = 0
	var excess = 0

	for y := initYear; y <= untilYear; y++ {
		var leap int

		if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
			utils.Debug("Year %d has leap-day", y)
			leap = 1
		} else {
			leap = 0
		}
		constants.DAYS_IN_MONTH[__FEBRUARY_KEY__] = 28 + leap

		for _, month := range constants.MONTHS_OF_YEAR {
			days := constants.DAYS_IN_MONTH[month]
			utils.Debug("Year %d| Month: %s | days %d", y, month, days)

			accumulatedRemainder += days % 7
			if accumulatedRemainder%7 == dayOfWeek {
				if y <= sinceYear {
					excess += 1
				}
				resultCount += 1
			}
		}
	}

	utils.Info("Problem0019 result: (%d - %d) => %d", resultCount, excess, resultCount-excess)

	return resultCount - excess
}
