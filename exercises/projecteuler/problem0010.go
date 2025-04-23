/**
 * @link Problem definition [[docs/projecteuler/problem0010.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/utils/log"
)

func Problem0010(bottom int, top int) int {

	var answer int

	var primes = []int{2}

	var i = bottom

	for i <= top {
		if helpers.IsPrime(i) {
			primes = append(primes, i)
			log.Debug("Prime found %d put in position: %d", i, len(primes))
		}
		i += 2
	}

	answer = helpers.Sum(primes)

	log.Info("primes count: %d", len(primes))
	log.Info("Sum of primes below ${top} is: ${sum(primes)}")

	log.Info("Problem0010 answer => Sum of primes below %d = %d", top, answer)

	return answer
}
