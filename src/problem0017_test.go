/**
 * Number letter counts
 *
 * https://projecteuler.net/problem=17
 *
 *
 * If the numbers 1 to 5 are written out in words:
 * one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19
 * letters used in total.
 *
 * If all the numbers from 1 to 1000 (one thousand) inclusive were written
 * out in words, how many letters would be used?
 *
 * NOTE: Do not count spaces or hyphens. For example, 342
 * (three hundred and forty-two) contains 23 letters and
 * 115 (one hundred and fifteen) contains 20 letters. The use of "and"
 * when writing out numbers is in compliance with British usage.
 */

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0017(t *testing.T) {

	const expectedSolution = 21124
	const inputInit = 1
	const inputLast = 1000

	testname := fmt.Sprintf("Problem0017(%d, %d) => %v \n", inputInit, inputLast, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0017(inputInit, inputLast)
		assert.Equal(t, expectedSolution, ans)
	})
}

func TestProblem0017ErrorBounds(t *testing.T) {

	const expectedSolution = -1
	const inputInit = 999
	const inputLast = 1001

	testname := fmt.Sprintf("Problem0017(%d, %d) => %v \n", inputInit, inputLast, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0017(inputInit, inputLast)
		assert.Equal(t, expectedSolution, ans)
	})
}
