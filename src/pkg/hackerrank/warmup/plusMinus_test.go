package hackerrank

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusMinus(t *testing.T) {

	expectedSolution := strings.Join([]string{"0.500000", "0.333333", "0.166667"}, "\n")
	input := []int{-4, 3, -9, 0, 4, 1}

	testname := fmt.Sprintf("PlusMinus(%d) => %v \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := PlusMinus(input)
		assert.Equal(t, expectedSolution, ans)
	})
}
