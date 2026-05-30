package helpers

func Product(numList []int) int {
	var result = 1

	if len(numList) == 0 {
		return 0
	}

	for i := range numList {
		result *= numList[i]
	}

	return result
}
