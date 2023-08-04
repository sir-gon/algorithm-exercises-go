// ////////////////////////////////////////////////////////////////////////////
// Result found:
//  Digits:
//      1071508607186267320948425049060001810561404811705
//      5336074437503883703510511249361224931983788156958
//      5812759467291755314682518714528569231404359845775
//      7469857480393456777482423098542107460506237114187
//      7954182153046474983581941267398767559165543946077
//      0629145711964776865421676604298316526243868372056
//      68069376
//  Sum: 1366
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"math/big"
	"strings"

	"gon.cl/algorithm-exercises/src/pkg/projecteuler/helpers"
	utils "gon.cl/algorithm-exercises/src/utils"
)

const __NUMERIC_BASE__ = 10
const __SPLIT_SEPARATOR__ = ""

func Problem0016(base int, exponent int) string {

	power := new(big.Int).Exp(big.NewInt(int64(base)), big.NewInt(int64(exponent)), nil)

	chars := strings.Split(power.Text(__NUMERIC_BASE__), __SPLIT_SEPARATOR__)

	answer := helpers.BigSumMany(chars)

	utils.Info("Problem0016 answer => %s", answer)

	return answer
}
