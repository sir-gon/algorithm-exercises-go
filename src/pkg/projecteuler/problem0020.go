/**
 * @link Problem definition [[docs/projecteuler/problem0020.md]]
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

	"gon.cl/algorithm-exercises/src/pkg/projecteuler/helpers"
	"gon.cl/algorithm-exercises/src/utils"
)

func Problem0020(last int) *big.Int {

	factorial := helpers.BigFactorial(last)

	utils.Info("Problem0020 answer => %v", factorial)

	return factorial
}
