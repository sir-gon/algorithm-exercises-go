package helpers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testIntMinOfManyLogMessage = "IntMinOfMany(%v); want %d"

func TestIntMinOfManyBasic(t *testing.T) {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedAnswer := 1
	expectedError := error(nil)

	testname := fmt.Sprintf(testIntMinOfManyLogMessage, inputList, expectedAnswer)
	t.Run(testname, func(t *testing.T) {
		ans, err := IntMinOfMany(inputList)
		assert.Equal(t, expectedAnswer, ans)
		assert.Equal(t, expectedError, err)
	})
}

func TestIntMinOfManyBorderCase(t *testing.T) {
	inputList := []int{}
	expectedAnswer := 0
	expectedError := errors.New("list is empty")

	testname := fmt.Sprintf(testIntMinOfManyLogMessage, inputList, expectedAnswer)
	t.Run(testname, func(t *testing.T) {
		ans, err := IntMinOfMany(inputList)
		assert.Equal(t, expectedAnswer, ans)
		assert.Equal(t, expectedError, err)
	})
}

func TestIntMaxOfManyBasic(t *testing.T) {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedAnswer := 10
	expectedError := error(nil)

	testname := fmt.Sprintf(testIntMinOfManyLogMessage, inputList, expectedAnswer)
	t.Run(testname, func(t *testing.T) {
		ans, err := IntMaxOfMany(inputList)
		assert.Equal(t, expectedAnswer, ans)
		assert.Equal(t, expectedError, err)
	})
}

func TestIntMaxOfManyBorderCase(t *testing.T) {
	inputList := []int{}
	expectedAnswer := 0
	expectedError := errors.New("list is empty")

	testname := fmt.Sprintf(testIntMinOfManyLogMessage, inputList, expectedAnswer)
	t.Run(testname, func(t *testing.T) {
		ans, err := IntMaxOfMany(inputList)
		assert.Equal(t, expectedAnswer, ans)
		assert.Equal(t, expectedError, err)
	})
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestIntMaxTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 1},
		{1, 0, 1},
		{2, -2, 2},
		{0, -1, 0},
		{-1, 0, 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMax(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

func BenchmarkIntMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMax(1, 2)
	}
}
