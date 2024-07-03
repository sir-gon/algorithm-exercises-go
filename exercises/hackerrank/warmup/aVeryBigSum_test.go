package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVeryBigSum(t *testing.T) {

	var input = []uint64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	const expectedSolution uint64 = 5000000015

	testname := fmt.Sprintf("solveMeFirst(%d) => %d \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := AVeryBigSum(input)
		assert.Equal(t, expectedSolution, ans)
	})
}
