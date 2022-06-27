/**
 * 10001st prime
 *
 * https://projecteuler.net/problem=7
 *
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
 * we can see that the 6th prime is 13.
 *
 * What is the 10 001st prime number?
 *
 */

package projecteuler

import (
	helpers "gon.cl/projecteuler.net/src/helpers"
	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0007(bottom int, top int) int {

	var primes []int

	i := bottom

	for doWhile := true; doWhile; doWhile = len(primes) < top {
		i += 1

		if helpers.IsPrime(i) {
			primes = append(primes, i)

			log.Debug("Prime found %d put in position: %d", i, len(primes))
		}
	}

	answer := primes[len(primes)-1]

	log.Info("Problem0007 asnwer => %d", answer)

	return answer
}
