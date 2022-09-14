package helpers

import (
	"sort"
)

func generateDivisorsOf(target int) []int {
	top := target
	var divs []int

	if target == 1 {
		divs = append(divs, 1)
		return divs
	}

	// fast divisors finding loop
	i := 1
	for i <= top {
		if target%i == 0 {
			divs = append(divs, i)
			divs = append(divs, top)
		}

		i++
		top = target / i
	}

	return divs
}

func Divisors(target int) []int {

	var divs []int = generateDivisorsOf(Abs(target))

	// sort terms
	sort.Ints(divs)

	return divs
}

func AreAmicables(a int, b int, _cache map[int]int) bool {

	if a == b || a <= 1 || b <= 1 {
		return false
	}

	if _cache[a] != 0 && _cache[b] != 0 {
		return _cache[a] == b && _cache[b] == a
	}

	caseA := Sum(Divisors(a))-a == b
	caseB := Sum(Divisors(b))-b == a

	return caseA && caseB
}
