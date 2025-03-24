/**
 * @link Problem definition [[docs/projecteuler/problem0023.md]]
 */

package projecteuler

import (
	"golang.org/x/exp/slices"
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	utils "gon.cl/algorithms/utils"
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

	var abundantNumberList = []int{}

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
