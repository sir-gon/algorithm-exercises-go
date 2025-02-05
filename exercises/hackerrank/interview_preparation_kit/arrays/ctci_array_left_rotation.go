package hackerrank

func rotLeftOne(a []int32) []int32 {
	var x int32
	var b []int32

	x, a = a[0], a[1:]
	b = append(b, a...)
	b = append(b, x)

	return b
}

func rotLeft(a []int32, d int32) []int32 {
	x := a[:]
	var i int32

	for i = 0; i < d; i++ {
		x = rotLeftOne(x)
	}
	return x
}

func RotLeft(a []int32, d int32) []int32 {
	return rotLeft(a, d)
}
