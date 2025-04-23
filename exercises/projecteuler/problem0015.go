/**
 * @link Problem definition [[docs/projecteuler/problem0015.md]]
 */

package projecteuler

import (
	helpers "gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/utils/log"
)

func Problem0015(gridSide int) int {

	var answer int

	vertexMatrix := helpers.Matrix(gridSide+1, gridSide+1, 1)

	for i := 1; i <= gridSide; i += 1 {
		for j := 1; j <= gridSide; j += 1 {
			upperParent := vertexMatrix[i-1][j]
			leftParent := vertexMatrix[i][j-1]

			vertexMatrix[i][j] = upperParent + leftParent
		}
	}
	log.Debug("Paths found: ${vertexMatrix[gridSide][gridSide]}")
	log.Debug("vertexMatrix: %v", vertexMatrix)

	log.Info("Problem0015 answer => %d", answer)

	return vertexMatrix[gridSide][gridSide]
}
