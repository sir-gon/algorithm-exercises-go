package projecteuler

import (
	"testing"
)

func TestProblem0002(t *testing.T) {

	solutionFound := 4613732
	top := 4000000

	ans := Problem0002(top)

	if ans != solutionFound {
		t.Errorf("Problem0002(%d) = %d; want %d", top, ans, solutionFound)
	}
}
