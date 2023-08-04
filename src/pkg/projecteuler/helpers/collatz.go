package helpers

func Collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func CollatzSequence(n int) []int {
	var sequence []int
	var c = n

	sequence = append(sequence, n)

	for c != 1 {
		c = Collatz(c)
		sequence = append(sequence, c)
		// console.log("sequence of ${i}: ${c}");
	}

	return sequence
}
