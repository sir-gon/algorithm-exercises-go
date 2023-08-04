package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixWithValueZero(t *testing.T) {
	want := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	const inputInitValue = 0
	const inputSizeM = 3
	const inputSizeN = 3

	testname := fmt.Sprintf("Matrix(%d, %d, %d) => %v", inputSizeM, inputSizeN, inputInitValue, want)
	t.Run(testname, func(t *testing.T) {
		ans := Matrix(inputSizeM, inputSizeN, inputInitValue)

		assert.Equal(t, want, ans)
	})
}

func TestMatrixWithValueOne(t *testing.T) {
	want := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	const inputInitValue = 1
	const inputSizeM = 3
	const inputSizeN = 3

	testname := fmt.Sprintf("Matrix(%d, %d, %d) => %v", inputSizeM, inputSizeN, inputInitValue, want)
	t.Run(testname, func(t *testing.T) {
		ans := Matrix(inputSizeM, inputSizeN, inputInitValue)

		assert.Equal(t, want, ans)
	})
}
