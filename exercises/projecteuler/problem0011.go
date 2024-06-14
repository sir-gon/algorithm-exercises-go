/**
 * @link Problem definition [[docs/projecteuler/problem0011.md]]
 */

package projecteuler

import (
	helpers "gon.cl/algorithms/exercises/projecteuler/helpers"
	utils "gon.cl/algorithms/utils"
)

func Problem0011(matrix [][]int, interval int) (int, bool) {
	if interval < 1 {
		return 0, true
	}

	var max int = 0

	var quadrantSize = interval
	var matrixLimit = len(matrix) - (interval - 1)

	for i := 0; i < matrixLimit; i += 1 {
		for j := 0; j < matrixLimit; j += 1 {
			utils.Debug("start point => i: %d, j: %d", i, j)

			// reset diagonals
			var diag1Acum = 1
			var diag2Acum = 1
			for k := 0; k < quadrantSize; k += 1 {
				utils.Debug("diag1 coordinate: (i, j) = (%d, %d)", i+k, j+k)
				utils.Debug("diag2 coordinate: (i, j) = (%d, %d)", i+k, j+(quadrantSize-1)-k)

				diag1Acum *= matrix[i+k][j+k]
				diag2Acum *= matrix[i+k][j+(quadrantSize-1)-k]

				max = helpers.IntMax(diag1Acum, max)
				max = helpers.IntMax(diag2Acum, max)

				// reset lines
				var verticalAcum = 1
				var horizontalAcum = 1
				for l := 0; l < quadrantSize; l += 1 {
					utils.Debug("vertical coordinate: (i, j) = (%d, %d)", i+k, j+l)
					utils.Debug("horizontal coordinate: (i, j) = (%d, %d)", i+l, j+k)

					verticalAcum *= matrix[i+k][j+l]
					horizontalAcum *= matrix[i+l][j+k]

					max = helpers.IntMax(verticalAcum, max)
					max = helpers.IntMax(horizontalAcum, max)
				}
			}
		}
	}

	utils.Info("Problem0011 max => %d", max)

	return max, false
}
