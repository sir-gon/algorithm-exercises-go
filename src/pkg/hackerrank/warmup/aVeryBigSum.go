/**
 * @link Problem definition [[docs/hackerrank/warmup/aVeryBigSum.md]]
 */

package hackerrank

import (
	utils "gon.cl/algorithm-exercises/src/utils"
)

func AVeryBigSum(ar []int) int {
	var result = 0

	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}

	utils.Info("aVeryBigSum answer => %d", result)
	return result
}
