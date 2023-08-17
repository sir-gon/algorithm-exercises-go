/**
 * @link Problem definition [[docs/projecteuler/problem0017.md]]
 */

package projecteuler

import (
	"math/big"
	"regexp"

	"gon.cl/algorithm-exercises/src/pkg/projecteuler/helpers"
	utils "gon.cl/algorithm-exercises/src/utils"
)

func Problem0017(init int, last int) int {

	var acum int

	var re = regexp.MustCompile("[^a-zA-Z0-9]+")

	for i := init; i <= last; i += 1 {

		word, errOutOfBounds := helpers.NumberToWord(*big.NewInt(int64(i)))
		if errOutOfBounds {
			return -1
		}

		replaced := re.ReplaceAllString(word, `$1`)

		acum += len(replaced)

		utils.Debug("acum: %d <= (%d) word: %s", acum, i, word)
	}

	utils.Info("Problem0017 answer => %d", acum)

	return acum
}
