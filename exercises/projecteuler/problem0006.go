/**
 * @link Problem definition [[docs/projecteuler/problem0006.md]]
 */

package projecteuler

import (
	"math"

	"gon.cl/algorithms/utils/log"
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
