package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0024SmallCase(t *testing.T) {

	expectedSolution := "ADCB"
	inputElements := "ABCD"
	const inputPermutationToFind = 6

	testname := fmt.Sprintf("Problem0024(%+v, %d) => %v \n", inputElements, inputPermutationToFind, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0024(inputElements, inputPermutationToFind)
		assert.Equal(t, expectedSolution, ans)
	})
}

func TestProblem0024(t *testing.T) {

	expectedSolution := "2783915460"
	inputElements := "0123456789"
	const inputPermutationToFind = 1000000

	testname := fmt.Sprintf("Problem0024(%+v, %d) => %v \n", inputElements, inputPermutationToFind, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0024(inputElements, inputPermutationToFind)
		assert.Equal(t, expectedSolution, ans)
	})
}
