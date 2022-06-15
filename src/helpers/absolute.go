package projecteuler

func Abs(value int) int {
	if value < 0 {
		return value * -1
	} else {
		return value
	}
}
