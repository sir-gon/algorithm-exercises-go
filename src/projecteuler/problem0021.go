/**
 *
 * Amicable numbers
 *
 * https://projecteuler.net/problem=21
 *
 * Let d(n) be defined as the sum of proper divisors of n
 * (numbers less than n which divide evenly into n).
 *
 * If d(a) = b and d(b) = a, where a ≠ b, then a and b are
 * an amicable pair and each of a and b are called amicable numbers.
 *
 * For example, the proper divisors of 220 are 1, 2, 4, 5, 10,
 * 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
 *
 * The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
 *
 * Evaluate the sum of all the amicable numbers under 10000.
 */

// ////////////////////////////////////////////////////////////////////////////
//
// Result found: 31626
//
// Amicable numbers found:
// amicableNumbers [
//      '220',  '284',
//      '1184', '1210',
//      '2620', '2924',
//      '5020', '5564',
//      '6232', '6368'
//    ]
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"gon.cl/algorithm-exercises/src/projecteuler/helpers"
	"gon.cl/algorithm-exercises/src/projecteuler/utils"
)

func Problem0021(_start int, _limit int) int {

	cache := map[int]int{}
	amicables := map[int]int{}

	for i := _start; i <= _limit; i++ {

		divs := helpers.Divisors(i)
		sum := helpers.Sum(divs) - i

		utils.Debug("Divisors of %d => %v => Sum = %d", i, divs, sum)

		cache[i] = sum
	}

	utils.Debug("Problem0021 cache => %d", cache)

	for a, b := range cache {
		if a != b && helpers.AreAmicables(a, b, cache) {
			amicables[a] = b
			amicables[b] = a
		}
	}

	var answer int

	utils.Info("Problem0021 amicables => %v", amicables)

	amicablesList := []int{}
	for value := range amicables {
		amicablesList = append(amicablesList, value)
	}

	answer = helpers.Sum(amicablesList)

	utils.Info("Problem0021 answer => %d", answer)

	return answer
}
