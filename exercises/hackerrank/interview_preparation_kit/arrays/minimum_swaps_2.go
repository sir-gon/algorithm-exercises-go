/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/arrays/minimum_swaps_2.md]]
 */

package hackerrank

func minimumSwaps(arr []int32) int32 {
	var size int32 = int32(len(arr))
	var indexedGroup = make(map[int32]int32, size)
	var swaps int32 = 0
	var index int32 = 0

	for key, value := range arr {
		indexedGroup[int32(key)] = value - 1
	}

	for index < size {
		if indexedGroup[index] == index {
			index += 1
		} else {
			var temp = indexedGroup[index]
			indexedGroup[index] = indexedGroup[temp]
			indexedGroup[temp] = temp
			swaps += 1
		}
	}

	return swaps
}

func MinimumSwaps(arr []int32) int32 {
	return minimumSwaps(arr)
}
