/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/greedy_algorithms/angry-children.md]]
 */

package hackerrank

import (
	"slices"
)

func maxMin(k int32, arr []int32) int32 {
	// Sort the array
	sortedlist := make([]int32, len(arr))
	copy(sortedlist, arr)
	slices.Sort(sortedlist)

	var result = sortedlist[len(arr)-1] - sortedlist[0]

	for i := range int32(len(arr)) - k + 1 {
		var tmin = sortedlist[i]
		var tmax = sortedlist[i+k-1]
		result = min(result, tmax-tmin)
	}

	return result
}

func MaxMin(k int32, arr []int32) int32 {
	return maxMin(k, arr)
}
