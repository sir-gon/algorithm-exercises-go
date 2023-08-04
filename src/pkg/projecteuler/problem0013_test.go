/**
 * Large sum
 *
 * https://projecteuler.net/problem=13
 *
 * Work out the first ten digits of the sum of the following
 * one-hundred 50-digit numbers.
 */

// ////////////////////////////////////////////////////////////////////////////
// See src/problem0013_data.go
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithm-exercises/src/pkg/projecteuler/data"
)

func TestProblem0013(t *testing.T) {

	expectedSolution := "5537376230"
	inputListOfBigNumbers := data.Problem0013Data

	testname := fmt.Sprintf("Problem0013(%s) => %v \n", inputListOfBigNumbers, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0013(inputListOfBigNumbers)
		assert.Equal(t, expectedSolution, ans)
	})
}
