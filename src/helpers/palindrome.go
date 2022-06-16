package helpers

import (
	"reflect"
)

// split integer into slice of single digits
func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func reverse(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)

	for i := len(b)/2 - 1; i >= 0; i-- {
		opp := len(b) - 1 - i
		b[i], b[opp] = b[opp], b[i]
	}

	return b
}

func IsPalindrome(n int) bool {
	digits := splitInt(n)
	reversed := reverse(digits)

	return reflect.DeepEqual(digits, reversed)
}
