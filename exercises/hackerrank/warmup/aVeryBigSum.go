/**
 * @link Problem definition [[docs/hackerrank/warmup/aVeryBigSum.md]]
 */

package hackerrank

import (
	utils "gon.cl/algorithms/utils"
)

func AVeryBigSum(ar []uint64) uint64 {
	var result uint64 = 0

	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}

	utils.Info("aVeryBigSum answer => %d", result)
	return result
}
