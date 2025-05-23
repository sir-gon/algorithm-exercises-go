// ////////////////////////////////////////////////////////////////////////////
// Result found:
//      Factorial of 100! =
//            93326215443944152681699238856266700490715968264381621
//            46859296389521759999322991560894146397615651828625369
//            7920827223758251185210916864000000000000000000000000
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils/log"
)

func TestProblem0020(t *testing.T) {

	input := 100

	expectedSolution := new(big.Int)
	expectedSolution, ok := expectedSolution.SetString("93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000", 10)

	if !ok {
		log.Error("Problem0020: expectedSolution.SetString() failed")
		return
	}

	testname := fmt.Sprintf("Problem0020(%d) => %v \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0020(input)
		assert.Equal(t, expectedSolution, ans)
	})

}
