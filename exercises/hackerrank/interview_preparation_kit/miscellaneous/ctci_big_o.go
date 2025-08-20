/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/miscellaneous/ctci-big-o.md]] # noqa
 */

package hackerrank

import "math"

const NOT_PRIME = "Not prime"
const PRIME = "Prime"

func primeFactor(n int32) *int32 {
	if n < 2 {
		return nil
	}

	divisor := n
	var maxPrimeFactor *int32 = nil

	var i int32 = 2
	for i <= int32(math.Sqrt(float64(divisor))) {
		if divisor%i == 0 {
			divisor = int32(divisor / i)
			maxPrimeFactor = &divisor
		} else {
			i += 1
		}
	}
	if maxPrimeFactor == nil {
		val := int32(n)
		return &val
	}

	return maxPrimeFactor
}

func primality(n int32) string {
	pf := primeFactor(n)
	if pf == nil || *pf != n {
		return NOT_PRIME
	}

	return PRIME
}

func Primality(n int32) string {
	return primality(n)
}
