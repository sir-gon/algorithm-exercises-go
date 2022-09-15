package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorsBasic(t *testing.T) {
	ans := Divisors(1)

	if len(ans) != 1 {

		t.Errorf("Divisors() = %d; want 1", ans)
	}
}

func TestDivisorsTableDriven(t *testing.T) {

	var tests = []struct {
		input int
		want  []int
	}{
		{input: 1, want: []int{1}},
		{input: 2, want: []int{1, 2}},
		{input: 10, want: []int{1, 2, 5, 10}},
		{input: 16, want: []int{1, 2, 4, 4, 8, 16}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Divisors(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := Divisors(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}

func TestNextPrimeFactorTableDriven(t *testing.T) {
	var tests = []struct {
		input  int
		answer Factor
	}{
		{input: 1, answer: Factor{factor: 1, carry: 1, cycles: 0}},
		{input: 2, answer: Factor{factor: 2, carry: 1, cycles: 1}},
		{input: 4, answer: Factor{factor: 2, carry: 2, cycles: 1}},
		{input: 9, answer: Factor{factor: 3, carry: 3, cycles: 2}},
		{input: 7, answer: Factor{factor: 7, carry: 1, cycles: 6}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Divisors(%d) => %v", tt.input, tt.answer)
		t.Run(testname, func(t *testing.T) {
			ans := NextPrimeFactor(tt.input)
			assert.Equal(t, tt.answer, ans)
		})
	}
}

func TestAreAmicableTableDriven(t *testing.T) {

	inputCache := map[int]int{}

	var tests = []struct {
		inputA int
		inputB int
		want   bool
	}{
		{inputA: 0, inputB: 1, want: false},
		{inputA: 1, inputB: 1, want: false},
		{inputA: 8, inputB: 8, want: false},
		{inputA: 220, inputB: 284, want: true},
		{inputA: 10, inputB: 15, want: false},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Divisors(%d, %d) => %v", tt.inputA, tt.inputB, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := AreAmicables(tt.inputA, tt.inputB, inputCache)
			assert.Equal(t, tt.want, ans)
		})
	}
}

func TestAreAmicableWithCache(t *testing.T) {

	inputA := 220
	inputB := 284
	inputCache := map[int]int{
		inputA: inputB,
		inputB: inputA,
	}

	expectedSolution := true

	ans := AreAmicables(inputA, inputB, inputCache)

	if ans != true {
		t.Errorf("AreAmicables(%d, %d, %v) = %t; want %t", inputA, inputB, inputCache, ans, expectedSolution)
	}
}
