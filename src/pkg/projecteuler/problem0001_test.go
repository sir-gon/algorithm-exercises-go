package projecteuler

import (
	"testing"
)

func TestProblem0001(t *testing.T) {

	solutionFound := 233168
	top := 1000

	ans := Problem0001(top)

	if ans != solutionFound {
		t.Errorf("Problem0001(%d) = %d; want %d", top, ans, solutionFound)
	}
}
