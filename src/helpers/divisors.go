package helpers

import (
	"sort"
)

func generateDivisorsOf(target int) []int {
	top := target
	var divs []int
	divs = append(divs, 1)

	if target != 1 {
		divs = append(divs, target)
	}

	// fast divisors finding loop
	for i := 2; i < top; i++ {
		top = target / i
		if target%i == 0 {
			divs = append(divs, i)

			if top != i {
				divs = append(divs, top)
			}
		}
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
