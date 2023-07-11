/**
 * Lexicographic Permutations
 *
 * https://projecteuler.net/problem=24
 *
 * A permutation is an ordered arrangement of objects.
 * For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
 * If all of the permutations are listed numerically or alphabetically,
 * we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
 *
 * 012   021   102   120   201   210
 *
 * What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
 */

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
