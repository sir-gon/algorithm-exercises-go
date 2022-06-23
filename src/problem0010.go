//go:build bruteforce
// +build bruteforce

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
	"gon.cl/projecteuler.net/src/helpers"
	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0010() int {

	var answer int

	const top = 2000000
	const bottom = 1

	var primes []int = []int{2}

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
