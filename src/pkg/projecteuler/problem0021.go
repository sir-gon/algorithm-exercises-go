/**
 * @link Problem definition [[docs/projecteuler/problem0021.md]]
 */
// ////////////////////////////////////////////////////////////////////////////
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
	"gon.cl/algorithm-exercises/src/pkg/projecteuler/helpers"
	"gon.cl/algorithm-exercises/src/utils"
)

func Problem0021(start int, limit int) int {

	cache := map[int]int{}
	amicables := map[int]int{}

	for i := start; i <= limit; i++ {

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
