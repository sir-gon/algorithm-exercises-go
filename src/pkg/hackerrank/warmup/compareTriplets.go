/**
 * @link Problem definition [[docs/hackerrank/warmup/compareTriplets.md]]
 */

package hackerrank

import (
	"errors"

	utils "gon.cl/algorithm-exercises/src/utils"
)

func CompareTriplets(a []int, b []int) ([]int, error) {

	if len(a) != len(b) {
		return nil, errors.New("empty input")
	}

	var awards = []int{0, 0}

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			awards[0] = awards[0] + 1
		}
		if a[i] < b[i] {
			awards[1] = awards[1] + 1
		}
	}

	utils.Info("SolveMeFirst answer => %v", awards)

	return awards, nil
}
