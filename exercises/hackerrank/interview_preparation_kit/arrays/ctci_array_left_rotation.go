package hackerrank

// @link Problem definition
// [[docs/hackerrank/interview_preparation_kit/arrays/ctci_array_left_rotation.md]]

func rotLeftOne(a []int32) []int32 {
	var x int32
	var b []int32

	x, a = a[0], a[1:]
	b = append(b, a...)
	b = append(b, x)

	return b
}

func rotLeft(a []int32, d int32) []int32 {
	// Sort the array
	x := make([]int32, len(a))
	copy(x, a)

	for range d {
		x = rotLeftOne(x)
	}
	return x
}

func RotLeft(a []int32, d int32) []int32 {
	return rotLeft(a, d)
}
