/**
 * @link Problem definition [[docs/hackerrank/warmup/diagonalDifference.md]]
 */

package hackerrank

import (
	"math"

	utils "gon.cl/algorithm-exercises/src/utils"
)

func DiagonalDifference(arr [][]int) int {
	diag1 := 0
	diag2 := 0
	last := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				diag1 += arr[i][j]
				diag2 += arr[last-i][j]
			}
		}
	}

	utils.Info("diag1 => %d", diag1)
	utils.Info("diag2 => %d", diag2)

	return int(math.Abs(float64(diag1 - diag2)))
}
