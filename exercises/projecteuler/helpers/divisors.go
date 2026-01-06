package helpers

import (
	"math"
	"slices"

	"gon.cl/algorithms/utils/log"
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

func Divisors(target int) []int {
	top := int(math.Abs(float64(target)))
	var divs []int

	divs = append(divs, int(1))

	if target == 1 {
		return divs
	}

	log.Debug("Find divisors of %d", target)

	// fast divisors finding loop
	i := int(2)
	for i <= top {
		top = int(target / i)
		var remainder = target % i

		if top > 2 && remainder == 0 {
			if i <= top {
				divs = append(divs, i)
			}

			if i < top {
				divs = append(divs, top)
			}
		}

		i += 1
	}

	divs = append(divs, target)
	log.Debug("collected divisors %x", divs)

	// sort divisors
	slices.Sort(divs)
	log.Debug("sorted divisors %x", divs)

	return divs
}

func ProperDivisors(target int) []int {
	var theDivisors = Divisors(target)

	theDivisors = theDivisors[:len(theDivisors)-1]

	return theDivisors
}

func NextPrimeFactor(target int) Factor {
	top := int(math.Abs(float64(target)))
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

func PrimeFactors(target int) (factors []int, cycles int) {
	factors = []int{}
	cycles = 0

	if target == 1 {
		factors = []int{1}
		cycles = 1
		return
	}

	var factor = target
	for factor != 1 {
		var partial = NextPrimeFactor(factor)
		cycles += partial.cycles

		factors = append(factors, partial.factor)
		factor = partial.carry
	}

	return
}

func AreAmicables(a, b int, cache map[int]int) bool {

	if a == b || a <= 1 || b <= 1 {
		return false
	}

	if cache[a] != 0 && cache[b] != 0 {
		return cache[a] == b && cache[b] == a
	}

	caseA := Sum(Divisors(a))-a == b
	caseB := Sum(Divisors(b))-b == a

	return caseA && caseB
}

const ___DIVISORS_ABUNDANT___ = "abundant"
const ___DIVISORS_PERFECT___ = "perfect"
const ___DIVISORS_DEFICIENT___ = "deficient"

type DIVISORS_ABUNDANCE string

const (
	DIVISORS_ABUNDANT  DIVISORS_ABUNDANCE = ___DIVISORS_ABUNDANT___
	DIVISORS_PERFECT   DIVISORS_ABUNDANCE = ___DIVISORS_PERFECT___
	DIVISORS_DEFICIENT DIVISORS_ABUNDANCE = ___DIVISORS_DEFICIENT___
)

func Abundance(target int) DIVISORS_ABUNDANCE {
	var divisors = ProperDivisors(target)

	divSum := 0
	for _, value := range divisors {
		divSum += value
	}

	if divSum > target {
		return DIVISORS_ABUNDANT
	}

	if divSum < target {
		return DIVISORS_DEFICIENT
	}

	return DIVISORS_PERFECT
}
