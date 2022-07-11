/**
 * Sum square difference
 *
 * https://projecteuler.net/problem=6
 *
 * The sum of the squares of the first ten natural numbers is,
 *
 * 1² * 2² * ... * 10² = 385
 *
 * The square of the sum of the first ten natural numbers is,
 *
 * (1 * 2 * ... * 10)² = 3025
 *
 * Hence the difference between the sum of the squares of the
 * first ten natural numbers and the square of the sum is
 * 3025 - 385.
 *
 * Find the difference between the sum of the squares of the
 * first one hundred natural numbers and the square of the sum.
 */

package projecteuler

import (
	"math"

	log "gon.cl/projecteuler.net/src/lib"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Problem0006() int {

	var answer int

	top := 100
	bottom := 1

	sumOfSquares := 0
	baseForSquareOfSum := 0
	squareOfSum := 0

	for i := bottom; i <= top; i++ {
		sumOfSquares += powInt(i, 2)
		baseForSquareOfSum += i
	}

	squareOfSum = powInt(baseForSquareOfSum, 2)

	log.Info("Sum of first %d squares = %d", top, sumOfSquares)
	log.Info("Base for Square Of Sum of first %d = %d", top, baseForSquareOfSum)
	log.Info("Square Of Sum of first %d = %d", top, squareOfSum)

	answer = squareOfSum - sumOfSquares

	log.Info("Problem0006 answer => %d - %d =  %d", squareOfSum, sumOfSquares, answer)

	return squareOfSum - sumOfSquares
}
