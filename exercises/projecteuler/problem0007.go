/**
 * @link Problem definition [[docs/projecteuler/problem0007.md]]
 */

package projecteuler

import (
	helpers "gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/utils/log"
)

func Problem0007(top int) int {

	var primes []int

	i := 0
	j := 2

	for doWhile := true; doWhile; doWhile = len(primes) < top {
		i += 1

		if helpers.IsPrime(j) {
			primes = append(primes, j)

			log.Debug("Prime found %d put in position: %d", j, len(primes))
		}

		j = 2*i + 1
	}

	answer := primes[len(primes)-1]

	cycles := i

	log.Info("Problem0007 answer => %d in %d cycles", answer, cycles)

	return answer
}
