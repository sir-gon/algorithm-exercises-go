/**
 * @link Problem definition
 * [[docs/hackerrank/interview_preparation_kit/greedy_algorithms/minimum-absolute-difference-in-an-array.md]]
 */

package hackerrank

import (
	"slices"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	// Sort the array
	sortedNums := make([]int32, len(arr))

	copy(sortedNums, arr)
	slices.Sort(sortedNums)

	// Find the minimum absolute difference
	result := int32(0)
	var resultEmpty = true

	for i := range len(sortedNums) - 1 {
		var aValue = sortedNums[i]
		var bValue = sortedNums[i+1]

		diff := aValue - bValue
		if diff < 0 {
			diff = -1 * diff
		}

		if resultEmpty {
			result = diff
			resultEmpty = false
		} else {
			result = min(result, diff)
		}
	}

	return result
}

func MinimumAbsoluteDifferenceInAnArray(arr []int32) int32 {
	return minimumAbsoluteDifference(arr)
}
