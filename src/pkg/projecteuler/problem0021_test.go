// ////////////////////////////////////////////////////////////////////////////
//
// Result found: 31626
//
// Amicable numbers found:
// amicableNumbers [
//      '220',  '284',
//      '1184', '1210',
//      '2620', '2924',
//      '5020', '5564',
//      '6232', '6368'
//    ]
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0021(t *testing.T) {

	expectedSolution := 31626
	inputStart := 1
	inputLimit := 10000

	testname := fmt.Sprintf("Problem0021(%d, %d) => %d \n", inputStart, inputLimit, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0021(inputStart, inputLimit)
		assert.Equal(t, expectedSolution, ans)
	})
}
