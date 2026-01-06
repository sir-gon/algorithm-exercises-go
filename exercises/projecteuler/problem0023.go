/**
 * @link Problem definition [[docs/projecteuler/problem0023.md]]
 */

package projecteuler

import (
	slices "slices"

	"gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/utils/log"
)

func Problem0023(underLimit, superLimit int) int {

	var abundantNumberList = []int{}

	// Produce a list of abundant numbers below limit
	for i := underLimit; i <= superLimit; i++ {
		var abundancyOf = helpers.Abundance(i)

		if abundancyOf == helpers.DIVISORS_ABUNDANT {
			abundantNumberList = append(abundantNumberList, i)
		}
	}

	log.Info("abundant nums list size => %d", len(abundantNumberList))
	log.Debug("abundant nums list result => %+v", abundantNumberList)

	var sumsAbundantNums = []int{}

	// Produce a list of sums of pair of abundant numbers below limit
	for i := range abundantNumberList {
		for j := 0; abundantNumberList[i]+abundantNumberList[j] <= superLimit &&
			j < len(abundantNumberList); j++ {

			var sum = abundantNumberList[i] + abundantNumberList[j]

			sumsAbundantNums = append(sumsAbundantNums, sum)
		}
	}

	log.Info("sumsAbundantNums size => %d", len(sumsAbundantNums))
	log.Debug("sumsAbundantNums result => %+v", sumsAbundantNums)

	// filter duplicates
	slices.Sort(sumsAbundantNums)
	sumsAbundantNums = slices.Compact(sumsAbundantNums)

	log.Info("filtered sumsAbundantNums size => %d", len(sumsAbundantNums))
	log.Debug("filtered sumsAbundantNums result => %+v", sumsAbundantNums)

	// All numbers below limit that not present in list of sums of pair of abundant numbers
	var found = []int{}

	for i := 1; i < superLimit; i++ {
		if !slices.Contains(sumsAbundantNums, i) {
			found = append(found, i)
		}
	}

	log.Info("found size %d", len(found))
	log.Debug("found result => %+v", found)

	result := 0
	for _, value := range found {
		result += value
	}

	log.Info("Problem0023 answer => %d", result)

	return result
}
