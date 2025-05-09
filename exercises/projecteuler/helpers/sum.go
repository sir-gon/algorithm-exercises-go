package helpers

func Sum(listOfNums []int) int {

	var sum int

	for i := range listOfNums {
		sum += listOfNums[i]
	}

	return sum
}
