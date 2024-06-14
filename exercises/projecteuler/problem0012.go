/**
 * @link Problem definition [[docs/projecteuler/problem0012.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	utils "gon.cl/algorithms/utils"
)

func Problem0012(top int) int {
	var amountOfDivisors = 0
	var triangular = 0
	var i = 1

	for amountOfDivisors < top {
		triangular += i
		var listOfDivisors = helpers.Divisors(triangular)
		amountOfDivisors = len(listOfDivisors)

		utils.Debug("Triangular number: %d has %d divisors", triangular, amountOfDivisors)

		i += 1
	}

	utils.Info("Problem0012 answer => %d", triangular)

	return triangular
}
