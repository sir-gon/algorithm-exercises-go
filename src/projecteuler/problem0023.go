/**
 * Non-abundant sums
 *
 * https://projecteuler.net/problem=23
 *
 * A perfect number is a number for which the sum of its proper
 * divisors is exactly equal to the number. For example, the sum
 * of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28,
 * which means that 28 is a perfect number.
 *
 * A number n is called deficient if the sum of its proper divisors is
 * less than n and it is called abundant if this sum exceeds n.
 *
 * As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the
 * smallest number that can be written as the sum of two abundant numbers is 24.
 * By mathematical analysis, it can be shown that all integers greater
 * than 28123 can be written as the sum of two abundant numbers
 * However, this upper limit cannot be reduced any further by analysis
 * even though it is known that the greatest number that cannot be expressed
 *  as the sum of two abundant numbers is less than this limit.
 *
 * Find the sum of all the positive integers which cannot be written as
 * the sum of two abundant numbers.
 */

package projecteuler

import (
	"golang.org/x/exp/slices"
	"gon.cl/algorithm-exercises/src/projecteuler/helpers"
	utils "gon.cl/algorithm-exercises/src/projecteuler/utils"
)

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Problem0023(underLimit int, superLimit int) int {

	var abundantNumberList []int = []int{}

	// Produce a list of abundant numbers below limit
	for i := underLimit; i <= superLimit; i++ {
		var abundancyOf = helpers.Abundance(i)

		if abundancyOf == helpers.DIVISORS_ABUNDANT {
			abundantNumberList = append(abundantNumberList, i)
		}
	}

	utils.Info("abundant nums list size => %d", len(abundantNumberList))
	utils.Debug("abundant nums list result => %+v", abundantNumberList)

	var sumsAbundantNums = []int{}

	// Produce a list of sums of pair of abundant numbers below limit
	for i := 0; i < len(abundantNumberList); i++ {
		for j := 0; abundantNumberList[i]+abundantNumberList[j] <= superLimit &&
			j < len(abundantNumberList); j++ {

			var sum = abundantNumberList[i] + abundantNumberList[j]

			sumsAbundantNums = append(sumsAbundantNums, sum)
		}
	}

	utils.Info("sumsAbundantNums size => %d", len(sumsAbundantNums))
	utils.Debug("sumsAbundantNums result => %+v", sumsAbundantNums)

	// filter duplicates
	slices.Sort(sumsAbundantNums)
	sumsAbundantNums = slices.Compact(sumsAbundantNums)

	utils.Info("filtered sumsAbundantNums size => %d", len(sumsAbundantNums))
	utils.Debug("filtered sumsAbundantNums result => %+v", sumsAbundantNums)

	// All numbers below limit that not present in list of sums of pair of abundant numbers
	var found = []int{}

	for i := 1; i < superLimit; i++ {
		if !Contains(sumsAbundantNums, i) {
			found = append(found, i)
		}
	}

	utils.Info("found size %d", len(found))
	utils.Debug("found result => %+v", found)

	result := 0
	for _, value := range found {
		result += value
	}

	utils.Info("Problem0023 answer => %d", result)

	return result
}
