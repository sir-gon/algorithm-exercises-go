package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0006(t *testing.T) {

	expectedSolution := 25164150

	testname := fmt.Sprintf("Problem0006() => %v \n", expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0006()
		assert.Equal(t, expectedSolution, ans)
	})
}
