/**
 * @link Problem definition [[docs/projecteuler/problem0012.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/utils/log"
)

func Problem0012(top int) int {
	var amountOfDivisors = 0
	var triangular = 0
	var i = 1

	for amountOfDivisors < top {
		triangular += i
		var listOfDivisors = helpers.Divisors(triangular)
		amountOfDivisors = len(listOfDivisors)

		log.Debug("Triangular number: %d has %d divisors", triangular, amountOfDivisors)

		i += 1
	}

	log.Info("Problem0012 answer => %d", triangular)

	return triangular
}
