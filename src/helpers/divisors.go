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
