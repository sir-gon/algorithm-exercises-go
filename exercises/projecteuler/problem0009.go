/**
 * @link Problem definition [[docs/projecteuler/problem0009.md]]
 */

package projecteuler

import (
	"math"

	utils "gon.cl/algorithms/utils"
)

func IsPythagoreanTriplet(a int, b int, c int) bool {
	return a < b && b < c && math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2)
}

func Problem0009() int {

	var answer int

	const balance = 1000

	var a = 1
	var b = 2
	var c = 997

	type triplet struct {
		a int
		b int
		c int
	}

	found := false
	var foundTriplet triplet

	for a < b && !found {
		b = a + 1

		for a < b && b < c && !found {

			utils.Debug("Pythagorean triplet? a = %d b = %d c = %d", a, b, c)

			if IsPythagoreanTriplet(a, b, c) {
				utils.Debug("FOUND: a = %d b = %d c = %d => %t", a, b, c, found)
				foundTriplet = triplet{a: a, b: b, c: c}
				found = true
			}

			b += 1
			c = balance - b - a
		}

		a += 1
	}

	utils.Info("FOUND: a = %d b = %d c = %d => %.0f + %.0f = %.0f",
		foundTriplet.a, foundTriplet.b, foundTriplet.c,
		math.Pow(float64(foundTriplet.a), 2),
		math.Pow(float64(foundTriplet.b), 2),
		math.Pow(float64(foundTriplet.c), 2))

	answer = foundTriplet.a * foundTriplet.b * foundTriplet.c
	utils.Info("Problem0009 answer => %d", answer)

	return answer
}
