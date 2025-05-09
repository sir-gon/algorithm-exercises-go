/**
 * @link Problem definition [[docs/hackerrank/warmup/simpleArraySum.md]]
 */

package hackerrank

import "gon.cl/algorithms/utils/log"

func SimpleArraySum(arr []int) int {
	acum := 0

	for i := range arr {
		acum += arr[i]
	}

	log.Info("SimpleArraySum answer => %d", acum)

	return acum
}
