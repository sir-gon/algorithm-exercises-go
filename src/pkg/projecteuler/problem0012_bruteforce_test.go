//go:build bruteforce
// +build bruteforce

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0012BruteForce(t *testing.T) {

	expectedSolution := 76576500
	inputTop := 500

	testname := fmt.Sprintf("Problem0012(%d) => %v \n", inputTop, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0012(inputTop)
		assert.Equal(t, expectedSolution, ans)
	})
}
