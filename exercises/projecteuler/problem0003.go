/**
 * @link Problem definition [[docs/projecteuler/problem0003.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	utils "gon.cl/algorithms/utils"
)

func Problem0003(top int) int {
	var maxPrimeFactor int

	divs := helpers.Divisors(top)

	utils.Info("Divisors(%d) = %v \n", top, divs)

	var i = len(divs) - 1

	for ok := true; ok; ok = i >= 0 && maxPrimeFactor == 0 {
		prime := helpers.IsPrime(divs[i])

		utils.Debug("%d is Prime? => %t \n", divs[i], prime)

		if prime {
			maxPrimeFactor = divs[i]
		}

		i -= 1
	}

	utils.Info("maxPrimeFactor is %d \n", maxPrimeFactor)

	return maxPrimeFactor
}
