package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0004(t *testing.T) {

	solutionFound := 906609
	bottom := 111
	top := 999

	testname := fmt.Sprintf("Problem0004(%d) => %v \n", top, solutionFound)
	t.Run(testname, func(t *testing.T) {
		ans := Problem0004(bottom, top)
		assert.Equal(t, solutionFound, ans)
	})
}
