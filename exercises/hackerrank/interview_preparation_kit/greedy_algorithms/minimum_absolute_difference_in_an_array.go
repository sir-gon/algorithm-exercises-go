/**
 * @link Problem definition
 * [[docs/hackerrank/interview_preparation_kit/greedy_algorithms/minimum-absolute-difference-in-an-array.md]]
 */

package hackerrank

import "slices"

func minimumAbsoluteDifference(arr []int32) int32 {
	// Sort the array
	arrCopy := make([]int32, len(arr))

	copy(arrCopy, arr)
	slices.Sort(arrCopy)

	// Find the minimum absolute difference
	result := int32(0)
	for i := range len(arrCopy) - 1 {
		diff := arrCopy[i+1] - arrCopy[i]
		if diff < result || result == 0 {
			result = diff
		}
	}
	return result
}

func MinimumAbsoluteDifferenceInAnArray(arr []int32) int32 {
	return minimumAbsoluteDifference(arr)
}
