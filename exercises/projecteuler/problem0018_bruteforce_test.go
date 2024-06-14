//go:build bruteforce
// +build bruteforce

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/exercises/projecteuler/data"
	"gon.cl/algorithms/exercises/projecteuler/lib"
)

func TestProblem0018BruteForce(t *testing.T) {

	inputTree, err := lib.BuildBNodeTreeWeigth(data.Problem0018_full_data)

	const expectedSolution = 1074
	const expectedErr = false

	testname := fmt.Sprintf("Problem0018() => %v \n", expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0018(inputTree)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, expectedSolution, ans)
	})
}
