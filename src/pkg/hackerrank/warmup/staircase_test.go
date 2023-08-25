package hackerrank

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaircase(t *testing.T) {

	input := 6
	expectedSolution := strings.Join([]string{
		"     #",
		"    ##",
		"   ###",
		"  ####",
		" #####",
		"######",
	}, "\n")

	testname := fmt.Sprintf("Staircase(%d) => %v \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Staircase(input)
		assert.Equal(t, expectedSolution, ans)
	})
}
