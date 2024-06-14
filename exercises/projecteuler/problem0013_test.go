// ////////////////////////////////////////////////////////////////////////////
// See data/problem0013_data.go
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/exercises/projecteuler/data"
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
