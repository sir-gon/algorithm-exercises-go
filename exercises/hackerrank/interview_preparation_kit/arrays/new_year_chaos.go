/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/arrays/new_year_chaos.md]]
 */

package hackerrank

import (
	"errors"
	"fmt"
)

//nolint:stylecheck // error strings should not be capitalized (ST1005)
const tooChaoticError = "Too chaotic"

func minimumBribesCalculate(q []int32) (int32, error) {

	var bribes int32
	var i int32 = 0
	var value int32

	for _, value = range q {
		position := i + 1

		if value-int32(position) > 2 {
			//nolint:stylecheck // error strings should not be capitalized (ST1005)
			return 0, errors.New(tooChaoticError)
		}

		var fragment []int32 = q[min(max(value-2, 0), int32(i)):i]

		for _, k := range fragment {
			if k > value {
				bribes += 1
			}
		}
		i += 1
	}

	return bribes, nil

}

func minimumBribes(q []int32) {
	result, err := minimumBribesCalculate(q)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(result)
}

func MinimumBribes(q []int32) {
	minimumBribes(q)
}

func MinimumBribesText(q []int32) string {
	result, err := minimumBribesCalculate(q)

	if err != nil {
		fmt.Println(err)

		return fmt.Sprint(err)
	}

	return fmt.Sprintf("%d", result)
}
