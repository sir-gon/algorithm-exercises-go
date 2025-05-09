/**
 * @link Problem definition [[docs/hackerrank/warmup/minimumiMaxSum.md]]
 */

package hackerrank

import (
	"errors"
	"fmt"
)

func MiniMaxSum(arr []int) (string, error) {
	if len(arr) == 0 {
		return "", errors.New("empty input")
	}

	sum := 0
	minimum := arr[0]
	maximum := arr[1]

	for i := range arr {
		num := arr[i]
		sum += num

		if num < minimum {
			minimum = num
		}

		if num > maximum {
			maximum = num
		}
	}

	return fmt.Sprintf("%d %d", sum-maximum, sum-minimum), nil
}
