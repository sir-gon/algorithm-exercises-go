/**
 * Power digit sum
 *
 * https://projecteuler.net/problem=16
 *
 *
 * 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
 *
 * What is the sum of the digits of the number 2^1000?
 */

// ////////////////////////////////////////////////////////////////////////////
// Result found:
//  Digits:
//      1071508607186267320948425049060001810561404811705
//      5336074437503883703510511249361224931983788156958
//      5812759467291755314682518714528569231404359845775
//      7469857480393456777482423098542107460506237114187
//      7954182153046474983581941267398767559165543946077
//      0629145711964776865421676604298316526243868372056
//      68069376
//  Sum: 1366
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0016(t *testing.T) {

	expectedSolution := "1366"
	inputBase := 2
	inputExponent := 1000

	testname := fmt.Sprintf("Problem0016(%d, %d) => %v \n", inputBase, inputExponent, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0016(inputBase, inputExponent)
		assert.Equal(t, expectedSolution, ans)
	})
}
