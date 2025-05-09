/**
 * @link Problem definition [[docs/hackerrank/warmup/diagonalDifference.md]]
 */

package hackerrank

import (
	"math"

	"gon.cl/algorithms/utils/log"
)

func DiagonalDifference(arr [][]int) int {
	diag1 := 0
	diag2 := 0
	last := len(arr) - 1

	for i := range arr {
		for j := range arr {
			if i == j {
				diag1 += arr[i][j]
				diag2 += arr[last-i][j]
			}
		}
	}

	log.Info("diag1 => %d", diag1)
	log.Info("diag2 => %d", diag2)

	return int(math.Abs(float64(diag1 - diag2)))
}
