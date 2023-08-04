package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0003(t *testing.T) {

	solutionFound := 6857
	top := 600851475143

	testname := fmt.Sprintf("Problem0003(%d) => %v", top, solutionFound)
	t.Run(testname, func(t *testing.T) {
		ans := Problem0003(top)
		assert.Equal(t, solutionFound, ans)
	})
}
