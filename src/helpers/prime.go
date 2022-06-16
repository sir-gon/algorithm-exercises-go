package helpers

// Prime checker
func IsPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
