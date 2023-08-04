/**
 * Summation of primes
 *
 * https://projecteuler.net/problem=10
 *
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 *
 * Find the sum of all the primes below two million.
 */

package projecteuler

import (
	"gon.cl/algorithm-exercises/src/projecteuler/helpers"
	utils "gon.cl/algorithm-exercises/src/projecteuler/utils"
)

func Problem0010(bottom int, top int) int {

	var answer int

	var primes []int = []int{2}

	var i = bottom

	for i <= top {
		if helpers.IsPrime(i) {
			primes = append(primes, i)
			utils.Debug("Prime found %d put in position: %d", i, len(primes))
		}
		i += 2
	}

	answer = helpers.Sum(primes)

	utils.Info("primes count: %d", len(primes))
	utils.Info("Sum of primes below ${top} is: ${sum(primes)}")

	utils.Info("Problem0010 answer => Sum of primes below %d = %d", top, answer)

	return answer
}
