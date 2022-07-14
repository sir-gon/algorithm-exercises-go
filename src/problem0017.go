/**
 * Number letter counts
 *
 * https://projecteuler.net/problem=17
 *
 *
 * If the numbers 1 to 5 are written out in words:
 * one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19
 * letters used in total.
 *
 * If all the numbers from 1 to 1000 (one thousand) inclusive were written
 * out in words, how many letters would be used?
 *
 * NOTE: Do not count spaces or hyphens. For example, 342
 * (three hundred and forty-two) contains 23 letters and
 * 115 (one hundred and fifteen) contains 20 letters. The use of "and"
 * when writing out numbers is in compliance with British usage.
 */

package projecteuler

import (
	"math/big"
	"regexp"

	"gon.cl/projecteuler.net/src/helpers"
	utils "gon.cl/projecteuler.net/src/utils"
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
