package helpers

import (
	"math/big"
)

const __NUMERIC_BASE__ = 10

func BigSumMany(strNumberArr []string) string {
	result := new(big.Int).SetInt64(int64(0))

	for i := range strNumberArr {
		var bignum, ok = new(big.Int).SetString(strNumberArr[i], 0)
		if !ok {
			return "0"
		}
		result = result.Add(result, bignum)
	}

	return result.Text(__NUMERIC_BASE__)
}

func BigFactorial(last int) *big.Int {

	const init = 1

	factorial := big.NewInt(int64(init))

	for i := init; i <= last; i += 1 {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}

	return factorial
}
