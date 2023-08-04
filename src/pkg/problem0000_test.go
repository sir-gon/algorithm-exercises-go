package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0000(t *testing.T) {

	expectedSolution := 0
	input := 0

	testname := fmt.Sprintf("Problem0000(%d) => %v \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0000()
		assert.Equal(t, expectedSolution, ans)
	})
}
