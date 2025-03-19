/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/2d_array.md]]
 */

package hackerrank

func arrayManipulation(n int32, queries [][]int32) int64 {
	// why adding 2?
	//   first slot to adjust 1-based index and
	//   last slot for storing accumSum result
	var LENGTH int32 = n + 2
	const InitialValue int64 = 0
	var maximum int64 = 0

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

		result[aStart] += int64(kValue)
		result[bEnd+1] -= int64(kValue)
	}

	var accumSum int64 = 0

	for _, value := range result {
		accumSum += value
		if accumSum > maximum {
			maximum = accumSum
		}
	}

	return maximum
}

func ArrayManipulation(n int32, queries [][]int32) int64 {
	return arrayManipulation(n, queries)
}
