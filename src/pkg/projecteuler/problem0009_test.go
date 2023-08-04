package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPythagoreanTriplet(t *testing.T) {

	var tests = []struct {
		a    int
		b    int
		c    int
		want bool
	}{
		{a: 0, b: 0, c: 0, want: false},
		{a: 1, b: 1, c: 1, want: false},
		{a: 2, b: 3, c: 4, want: false},
		{a: 3, b: 4, c: 5, want: true},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("IsPythagoreanTriplet(%d, %d, %d) => %v", tt.a, tt.b, tt.c, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := IsPythagoreanTriplet(tt.a, tt.b, tt.c)
			assert.Equal(t, tt.want, ans)
		})
	}

}

func TestProblem0009(t *testing.T) {

	const a = 200
	const b = 375
	const c = 425

	expectedSolution := a * b * c // 31875000

	testname := fmt.Sprintf("Problem0009() => %v \n", expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0009()
		assert.Equal(t, expectedSolution, ans)
	})
}
