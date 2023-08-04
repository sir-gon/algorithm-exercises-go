package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithm-exercises/src/pkg/projecteuler/data"
	"gon.cl/algorithm-exercises/src/pkg/projecteuler/lib"
)

func TestProblem0018(t *testing.T) {

	inputTree, err := lib.BuildBNodeTreeWeigth(data.Problem0018_small_data)

	const expectedSolution = 23
	const expectedErr = false

	testname := fmt.Sprintf("Problem0018() => %v \n", expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0018(inputTree)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, expectedSolution, ans)
	})
}
