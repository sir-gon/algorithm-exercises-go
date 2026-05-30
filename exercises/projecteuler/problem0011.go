/**
 * @link Problem definition [[docs/projecteuler/problem0011.md]]
 */

package projecteuler

import (
	helpers "gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/utils/log"
)

func Problem0011(matrix [][]int, interval int) (int, bool) {
	if interval < 1 {
		return 0, true
	}

	var greatest = 0

	var quadrantSize = interval
	var matrixLimit = len(matrix) - (interval - 1)

	for i := range matrixLimit {
		for j := range matrixLimit {
			log.Debug("start point => i: %d, j: %d", i, j)

			// reset diagonals
			var diag1Acum = 1
			var diag2Acum = 1
			for k := range quadrantSize {
				log.Debug("diag1 coordinate: (i, j) = (%d, %d)", i+k, j+k)
				log.Debug("diag2 coordinate: (i, j) = (%d, %d)", i+k, j+(quadrantSize-1)-k)

				diag1Acum *= matrix[i+k][j+k]
				diag2Acum *= matrix[i+k][j+(quadrantSize-1)-k]

				greatest = helpers.IntMax(diag1Acum, greatest)
				greatest = helpers.IntMax(diag2Acum, greatest)

				// reset lines
				var verticalAcum = 1
				var horizontalAcum = 1
				for l := range quadrantSize {
					log.Debug("vertical coordinate: (i, j) = (%d, %d)", i+k, j+l)
					log.Debug("horizontal coordinate: (i, j) = (%d, %d)", i+l, j+k)

					verticalAcum *= matrix[i+k][j+l]
					horizontalAcum *= matrix[i+l][j+k]

					greatest = helpers.IntMax(verticalAcum, greatest)
					greatest = helpers.IntMax(horizontalAcum, greatest)
				}
			}
		}
	}

	log.Info("Problem0011 greatest => %d", greatest)

	return greatest, false
}
