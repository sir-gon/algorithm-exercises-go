/**
 * Lattice paths
 *
 * https://projecteuler.net/problem=15
 *
 *
 * Starting in the top left corner of a 2×2 grid, and only being able to move
 * to the right and down, there are exactly 6 routes to the bottom right corner.
 *
 * How many such routes are there through a 20×20 grid?
 */

package projecteuler

import (
	helpers "gon.cl/projecteuler.net/src/helpers"
	log "gon.cl/projecteuler.net/src/lib"
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
