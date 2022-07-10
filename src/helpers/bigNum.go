package helpers

import "math/big"

const __NUMERIC_BASE__ = 10

func BigSumMany(strNumberArr []string) string {
	result := new(big.Int).SetInt64(int64(0))

	for i := 0; i < len(strNumberArr); i += 1 {
		var bignum, ok = new(big.Int).SetString(strNumberArr[i], 0)
		if !ok {
			return "0"
		}
		result = result.Add(result, bignum)
	}

	return result.Text(__NUMERIC_BASE__)
}
