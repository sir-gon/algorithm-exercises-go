//go:build bruteforce
// +build bruteforce

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0010BruteForce(t *testing.T) {

	const expectedSolution = 142913828922
	const bottom = 1
	const top = 2000000

	testname := fmt.Sprintf("Problem0010(%d, %d) => %v \n", bottom, top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0010(bottom, top)
		assert.Equal(t, expectedSolution, ans)
	})
}
