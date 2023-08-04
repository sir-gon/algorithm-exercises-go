// ////////////////////////////////////////////////////////////////////////////
// See:
//    - src/data/p022_listOfNames.txt
//    - src/data/p022_listOfNames.json
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithm-exercises/src/pkg/projecteuler/data"
)

func TestProblem0022(t *testing.T) {

	expectedSolution := 871198282

	testname := fmt.Sprintf("Problem0022(data.Problem0022_data) => %v \n", expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0022(data.Problem0022_data)
		assert.Equal(t, expectedSolution, ans)
	})
}
