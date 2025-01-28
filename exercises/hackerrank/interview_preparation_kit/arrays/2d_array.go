/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/2d_array.md]]
 */

package hackerrank

func getHourGlass(arr [][]int32, positionX int32, positionY int32) []int32 {
	result := []int32{}

	// top
	result = append(result, arr[positionX-1][positionY-1])
	result = append(result, arr[positionX-1][positionY])
	result = append(result, arr[positionX-1][positionY+1])

	// middle
	result = append(result, arr[positionX][positionY])

	// bottom
	result = append(result, arr[positionX+1][positionY-1])
	result = append(result, arr[positionX+1][positionY])
	result = append(result, arr[positionX+1][positionY+1])

	return result
}

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	matrixSize := len(arr)

	matrixStartIndex := 1
	matrixEndIndex := matrixSize - 2

	var maxHourglassSum int32 = 0

	for i := int32(matrixStartIndex); i <= int32(matrixEndIndex); i++ {
		for j := int32(matrixStartIndex); j <= int32(matrixEndIndex); j++ {
			var currentHourglassSum int32 = 0
			currentHourglass := getHourGlass(arr, i, j)
			for k := 0; k < len(currentHourglass); k++ {
				currentHourglassSum += currentHourglass[k]
			}

			if currentHourglassSum > maxHourglassSum {
				maxHourglassSum = currentHourglassSum
			}
		}
	}

	return maxHourglassSum
}

func HourglassSum(arr [][]int32) int32 {
	return hourglassSum(arr)
}
