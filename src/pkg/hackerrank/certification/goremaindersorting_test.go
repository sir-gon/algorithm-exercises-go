package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveMeFirst(t *testing.T) {

	a := []string{
		"Colorado",
		"Utah",
		"Montana",
		"Wisconsin",
		"Oregon",
		"Maine",
	}
	expectedSolution := []string{
		"Oregon",
		"Wisconsin",
		"Montana",
		"Utah",
		"Colorado",
		"Maine",
	}

	testname := fmt.Sprintf("RemainderSorting(%v) => %v \n", a, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := RemainderSorting(a)
		assert.Equal(t, expectedSolution, ans)
	})
}
