/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/arrays/crush.md]]
 */

package hackerrank

import "slices"

func arrayManipulationBruteForce(n int32, queries [][]int32) int64 {
	var LENGTH int32 = n + 1
	const InitialValue int64 = 0

	result := make([]int64, LENGTH)
	for i := 0; int32(i) < LENGTH; i++ {
		result[i] = InitialValue
	}

	var aStart int32
	var bEnd int32
	var kValue int32

	for _, query_row := range queries {
		aStart = query_row[0]
		bEnd = query_row[1]
		kValue = query_row[2]

		for i := aStart; i <= bEnd; i++ {
			result[i] += int64(kValue)
		}
	}

	return int64(slices.Max(result))

}

func ArrayManipulationBruteForce(n int32, queries [][]int32) int64 {
	return arrayManipulationBruteForce(n, queries)
}
