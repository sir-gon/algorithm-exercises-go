/**
 * @link Problem definition [[docs/hackerrank/warmup/simpleArraySum.md]]
 */

package hackerrank

import (
	utils "gon.cl/algorithms/utils"
)

func SimpleArraySum(arr []int) int {
	acum := 0

	for i := 0; i < len(arr); i++ {
		acum += arr[i]
	}

	utils.Info("SimpleArraySum answer => %d", acum)

	return acum
}
