package helpers

func Sum(listOfNums []int) int {

	var sum int

	for i := 0; i < len(listOfNums); i++ {
		sum += listOfNums[i]
	}

	return sum
}
