/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/greedy_algorithms/greedy-florist.md]]
 * @see Solution notes [[docs/hackerrank/interview_preparation_kit/greedy_algorithms/greedy-florist-solution-notes.md]]
 */

package hackerrank

import (
	"slices"
)

func getMinimumCost(k int32, c []int32) int32 {
	// Sort the array
	flowers := make([]int32, len(c))
	copy(flowers, c)
	slices.Sort(flowers)
	slices.Reverse(flowers)

	total := int32(0)
	k_customers := int(k)

	for i, flower_cost := range flowers {
		var position = int32(i / k_customers)
		total += int32(position+1) * int32(flower_cost)
	}

	return total
}

func GetMinimumCost(k int32, c []int32) int32 {
	return getMinimumCost(k, c)
}
