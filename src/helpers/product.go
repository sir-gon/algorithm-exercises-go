package helpers

func Product(numList []int) int {
	var result int = 1
	var i int = 0

	if len(numList) == 0 {
		return 0
	}

	for i = 0; i < len(numList); i += 1 {
		result *= numList[i]
	}

	return result
}
