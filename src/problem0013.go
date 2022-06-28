/**
 * Large sum
 *
 * https://projecteuler.net/problem=13
 *
 * Work out the first ten digits of the sum of the following
 * one-hundred 50-digit numbers.
 */

// ////////////////////////////////////////////////////////////////////////////
// See src/problem0013_data.go
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"math/big"

	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0013(inputListOfBigNumbers []string) string {

	const _BASE_ = 10
	const top = 10000000000
	answer := new(big.Int)
	bigTop := new(big.Int).SetInt64((top))

	var listOfBigNumbers []*big.Int

	var i int
	for i = 0; i < len(inputListOfBigNumbers); i += 1 {
		var bignum, ok = new(big.Int).SetString(inputListOfBigNumbers[i], 0)
		listOfBigNumbers = append(listOfBigNumbers, bignum)

		log.Debug("new bigNumber: %s | %t", bignum.Text(_BASE_), ok)

		answer = answer.Add(answer, bignum)
	}

	log.Debug("listOfBigNumbers: %v", listOfBigNumbers)

	for answer.Cmp(bigTop) >= 1 {
		answer = answer.Div(answer, big.NewInt(_BASE_))

		log.Debug("Answer reduction: %s", answer.String())
	}

	log.Info("Problem0013 answer => %v", answer)

	return answer.Text(_BASE_)
}
