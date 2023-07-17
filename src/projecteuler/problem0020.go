/**
 * Factorial digit sum
 *
 * https://projecteuler.net/problem=20
 *
 * n! means n × (n − 1) × ... × 3 × 2 × 1
 *
 * For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
 * and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
 *
 * Find the sum of the digits in the number 100!
 */

// ////////////////////////////////////////////////////////////////////////////
// Result found:
//      Factorial of 100! =
//            93326215443944152681699238856266700490715968264381621
//            46859296389521759999322991560894146397615651828625369
//            7920827223758251185210916864000000000000000000000000
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"math/big"

	"gon.cl/projecteuler.net/src/projecteuler/helpers"
	"gon.cl/projecteuler.net/src/projecteuler/utils"
)

func Problem0020(_last int) *big.Int {

	factorial := helpers.BigFactorial(_last)

	utils.Info("Problem0020 answer => %v", factorial)

	return factorial
}
