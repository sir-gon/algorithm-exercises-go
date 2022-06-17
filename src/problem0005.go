/**
* Smallest multiple
*
* https://projecteuler.net/problem=5
*
*
*  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
*
*  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*
* ////////////////////////////////////////////////////////////////////////////
* NOTE: I think a better way to solve problem 0005.
* If I got prime numbers in beetwen [1,20], can I test them as
* prime factors, testing a combinatorial exponents that satisfies
* the condition that the number found is divisible by all
* the numbers in beetwen [1,20],
* Instead, this solution is brute force.
*
/ ////////////////////////////////////////////////////////////////////////////
* Solution found by BRUTE FORCE running for a LONG TIME this routine
*
* FOUND: 232792560 divisible by any element beetwen 1 and 20
/ ////////////////////////////////////////////////////////////////////////////
*/

package projecteuler

import (
	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0005(bottom int, top int, startFrom int) int {

	var answer int
	found := false

	var fail bool
	var i int
	test := startFrom

	for doWhile := true; doWhile; doWhile = !found {

		i = 2
		// fail = false

		for doWhile := true; doWhile; doWhile = (i <= top && !fail) {

			fail = test%i != 0

			log.Debug("%d divisible by %d (remainder: %d => %t)", test, i, test%i, fail)

			i += 1
		}

		if !fail {
			answer = test
			found = true
		}

		fail = false
		test += 1
	}

	log.Info("Problem0005 answer => %d divisible by any element beetwen %d and %d", answer, bottom, top)

	return answer
}
