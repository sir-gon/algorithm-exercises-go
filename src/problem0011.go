/**
 *
 * https://projecteuler.net/problem=XX
 *
 *
 */

package projecteuler

import (
	"gon.cl/projecteuler.net/src/helpers"
	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0011(matrix [][]int, step int) (int, bool) {

	// const bottom = 0
	var top = len(matrix)
	// const step = 4

	if step < 1 {
		return 0, true
	}

	var max int = 0
	var acum int = 0

	var i, j, k int

	for i = 0; i < top; i += 1 {
		for j = 0; j < top; j += 1 {
			acum = 1

			if i < top-(step-1) && j < top {
				log.Debug("---- VERTICAL ------------------------------------------")
				// vertical
				for k = 0; k < step; k++ {
					log.Debug("row: i %d, column: %d, step %d => %d", i+k, j, k, matrix[i+k][j])

					acum *= matrix[i+k][j]
				}

				max = helpers.IntMax(acum, max)
			}

			acum = 1
			if i < top && j < top-(step-1) {
				log.Debug("---- HORIZONTAL ----------------------------------------")
				// horizontal
				for k = 0; k < step; k++ {
					log.Debug("row: i ${i}, column: ${j + k} => ${matrix[i][j + k]}")
					acum *= matrix[i][j+k]
				}

				max = helpers.IntMax(acum, max)
			}

			acum = 1
			if i+(step-1) < top && j+(step-1) < top {
				// diagonal top left -> bottom right
				log.Debug("---- DIAG \\ ---------------------------------------------")
				for k = 0; k < step; k++ {
					log.Debug("diag: (${i + k}, ${j + k}) => ${matrix[i + k][j + k]}")
					acum *= matrix[i+k][j+k]
				}

				max = helpers.IntMax(acum, max)
			}

			acum = 1
			if i+(step-1) < top && j+(step-1) < top {
				// diagonal top rigth -> bottom left
				log.Debug("---- DIAG / ---------------------------------------------")
				for k = 0; k < step; k++ {
					log.Debug("diag: (${i + k}, ${j + (step - 1) - k}) => ${matrix[i + k][j + (step - 1) - k]}")
					acum *= matrix[i+k][j+(step-1)-k]
				}

				max = helpers.IntMax(acum, max)
			}

		}
	}

	log.Info("Problem0011 max => %d", max)

	return max, false
}
