/**
 * @link Problem definition [[docs/hackerrank/warmup/miniMaxSum.md]]
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
	min := arr[0]
	max := arr[1]

	for i := range arr {
		num := arr[i]
		sum += num

		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return fmt.Sprintf("%d %d", sum-max, sum-min), nil
}
