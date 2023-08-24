/**
 * @link Problem definition [[docs/projecteuler/problem0001.md]]
 */

package projecteuler

import (
	utils "gon.cl/algorithm-exercises/src/utils"
)

func Problem0001(top int) int {
	total := 0

	for i := 0; i < top; i++ {
		if i%3 == 0 || i%5 == 0 {

			total += i

			utils.Debug("Problem0001: STEP: %d => TOTAL: %d", i, total)
		}
	}

	utils.Info("Problem0001 result: %d", total)

	return total
}
