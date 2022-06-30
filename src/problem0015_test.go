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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0015(t *testing.T) {

	expectedSolution := 137846528820
	inputGridSize := 20

	testname := fmt.Sprintf("Problem0015(%d) => %v \n", inputGridSize, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0015(inputGridSize)
		assert.Equal(t, expectedSolution, ans)
	})
}
