package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleArraySum(t *testing.T) {

	expectedSolution := 31
	input := []int{1, 2, 3, 4, 10, 11}

	testname := fmt.Sprintf("SimpleArraySum(%v) => %d \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := SimpleArraySum(input)
		assert.Equal(t, expectedSolution, ans)
	})
}
