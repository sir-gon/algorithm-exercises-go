/**
 * @link Problem definition [[docs/projecteuler/problem0002.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/utils/log"
)

func Problem0002(top int) int {

	last1 := 1
	last2 := 0
	evenSum := 0

	i := 0
	fibo := 0

	for ok := true; ok; ok = fibo < top {
		fibo = last2 + last1

		log.Debug("Fibonacci (%d) = %d", i, fibo)

		// acumulate sum of event terms
		if fibo%2 == 0 {
			evenSum += fibo
		}

		// next keys in loop:
		last2 = last1
		last1 = fibo
		i += 1
	}

	log.Info("Problem0002 result: %d", evenSum)

	return evenSum
}
