/**
 * @link Problem definition [[docs/hackerrank/warmup/aVeryBigSum.md]]
 */

package hackerrank

import "gon.cl/algorithms/utils/log"

func AVeryBigSum(ar []uint64) uint64 {
	var result uint64 = 0

	for i := range ar {
		result += ar[i]
	}

	log.Info("aVeryBigSum answer => %d", result)
	return result
}
