package projecteuler

import (
	utils "gon.cl/algorithm-exercises/src/utils"
)

func Problem0002(top int) int {

	last1 := 1
	last2 := 0
	evenSum := 0

	i := 0
	fibo := 0

	for ok := true; ok; ok = fibo < top {
		fibo = last2 + last1

		utils.Debug("Fibonacci (%d) = %d", i, fibo)

		// acumulate sum of event terms
		if fibo%2 == 0 {
			evenSum += fibo
		}

		// next keys in loop:
		last2 = last1
		last1 = fibo
		i += 1
	}

	utils.Info("Problem0002 result: %d", evenSum)

	return evenSum
}
