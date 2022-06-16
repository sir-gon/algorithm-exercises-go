/**
 *
 * https://projecteuler.net/problem=XX
 *
 *
 */

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0000Template(t *testing.T) {

	expectedSolution := 0
	input := 0

	testname := fmt.Sprintf("Problem0000Template(%d) => %v \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0000Template()
		assert.Equal(t, expectedSolution, ans)
	})
}
