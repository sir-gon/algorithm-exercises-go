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
	helpers "gon.cl/algorithm-exercises/src/projecteuler/helpers"
	utils "gon.cl/algorithm-exercises/src/projecteuler/utils"
)

func Problem0007(top int) int {

	var primes []int

	i := 0
	j := 2

	for doWhile := true; doWhile; doWhile = len(primes) < top {
		i += 1

		if helpers.IsPrime(j) {
			primes = append(primes, j)

			utils.Debug("Prime found %d put in position: %d", j, len(primes))
		}

		j = 2*i + 1
	}

	answer := primes[len(primes)-1]

	cycles := i

	utils.Info("Problem0007 answer => %d in %d cycles", answer, cycles)

	return answer
}
