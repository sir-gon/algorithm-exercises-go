package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveMeFirst(t *testing.T) {

	expectedSolution := 0
	a := 0
	b := 0

	testname := fmt.Sprintf("solveMeFirst(%d, %d) => %v \n", a, b, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := SolveMeFirst(a, b)
		assert.Equal(t, expectedSolution, ans)
	})
}
