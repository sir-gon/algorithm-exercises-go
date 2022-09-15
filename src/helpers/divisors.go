package helpers

import (
	"math"
	"sort"
)

type Factor struct {
	factor int
	carry  int
	cycles int
}

type Factors struct {
	factors []int
	cycles  int
}

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

func NextPrimeFactor(_target int) Factor {
	top := int(math.Abs(float64(_target)))
	cycles := 0

	if top == 1 {
		return Factor{
			factor: 1,
			carry:  1,
			cycles: cycles,
		}
	}

	var i = 2
	for i < top {
		cycles += 1
		if top%i == 0 {
			return Factor{
				factor: i,
				carry:  top / i,
				cycles: cycles,
			}
		}
		i += 1
	}

	return Factor{
		factor: i,
		carry:  top / i,
		cycles: cycles + 1,
	}
}

func PrimeFactors(target int) Factors {
	var factors = Factors{
		factors: []int{},
		cycles:  0,
	}

	if target == 1 {
		return Factors{
			factors: []int{1},
			cycles:  1,
		}
	}

	var factor = target
	for factor != 1 {
		var partial = NextPrimeFactor(factor)
		factors.cycles += partial.cycles

		factors.factors = append(factors.factors, partial.factor)
		factor = partial.carry
	}

	return factors
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
