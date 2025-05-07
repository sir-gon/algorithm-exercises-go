// @link Problem definition [[docs/hackerrank/interview_preparation_kit/greedy_algorithms/luck-balance.md]] // noqa

package hackerrank

import (
	"cmp"
	"slices"
)

type Contest struct {
	luck      int32
	important int32
}

func luckBalance(k int32, contests [][]int32) int32 {
	var important_contests = []Contest{}
	var nonimportant_contests = []Contest{}

	for _, contest := range contests {
		var contest = Contest{
			luck:      contest[0],
			important: contest[1],
		}

		if contest.important == 1 {

			important_contests = append(important_contests, contest)
		} else {
			nonimportant_contests = append(nonimportant_contests, contest)
		}
	}

	slices.SortFunc(
		important_contests,
		func(a, b Contest) int {
			return cmp.Or(
				-cmp.Compare(a.important, b.important),
				-cmp.Compare(a.luck, b.luck),
			)
		})

	var total int32 = 0
	var size = int32(len(important_contests))
	var cut = min(k, int32(size))

	for i := 0; int32(i) < cut; i++ {
		total += important_contests[i].luck
	}

	for i := cut; i < size; i++ {
		total -= important_contests[i].luck
	}

	for _, contest := range nonimportant_contests {
		total += contest.luck
	}

	return total
}
