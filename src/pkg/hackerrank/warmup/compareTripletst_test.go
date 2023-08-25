package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareTripletsBorderCase(t *testing.T) {
	expectedSolution := []int{1, 1}
	a := []int{1}
	b := []int{2, 3}

	testname := fmt.Sprintf("CompareTriplets(%d, %d) => %v \n", a, b, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		_, err := CompareTriplets(a, b)
		assert.Error(t, err)
	})
}

func TestCompareTriplets(t *testing.T) {

	expectedSolution := []int{1, 1}
	a := []int{5, 6, 7}
	b := []int{3, 6, 10}

	testname := fmt.Sprintf("CompareTriplets(%d, %d) => %v \n", a, b, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans, _ := CompareTriplets(a, b)
		assert.Equal(t, expectedSolution, ans)
	})
}
