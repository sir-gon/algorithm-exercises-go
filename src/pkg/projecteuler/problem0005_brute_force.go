package projecteuler

import (
	utils "gon.cl/algorithm-exercises/src/utils"
)

func Problem0005BruteForce(bottom int, top int, startFrom int) int {

	var answer int
	found := false

	var fail bool
	var i int
	test := startFrom

	for doWhile := true; doWhile; doWhile = !found {

		i = 2
		// fail = false

		for doWhile := true; doWhile; doWhile = (i <= top && !fail) {

			fail = test%i != 0

			utils.Debug("%d divisible by %d (remainder: %d => %t)", test, i, test%i, fail)

			i += 1
		}

		if !fail {
			answer = test
			found = true
		}

		fail = false
		test += 1
	}

	utils.Info("Problem0005 answer => %d divisible by any element beetwen %d and %d", answer, bottom, top)

	return answer
}
