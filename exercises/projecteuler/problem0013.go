/**
 * @link Problem definition [[docs/projecteuler/problem0013.md]]
 * @see [[src/pkg/projecteuler/data/problem0013_data.go]]
 */

package projecteuler

import (
	"math/big"

	"gon.cl/algorithms/utils/log"
)

func Problem0013(inputListOfBigNumbers []string) string {

	const BASE = 10
	const top = 10000000000
	answer := new(big.Int)
	bigTop := new(big.Int).SetInt64((top))

	var listOfBigNumbers []*big.Int

	for i := 0; i < len(inputListOfBigNumbers); i += 1 {
		var bignum, ok = new(big.Int).SetString(inputListOfBigNumbers[i], 0)
		listOfBigNumbers = append(listOfBigNumbers, bignum)

		log.Debug("new bigNumber: %s | %t", bignum.Text(BASE), ok)

		answer = answer.Add(answer, bignum)
	}

	log.Debug("listOfBigNumbers: %v", listOfBigNumbers)

	for answer.Cmp(bigTop) >= 1 {
		answer = answer.Div(answer, big.NewInt(BASE))

		log.Debug("Answer reduction: %s", answer.String())
	}

	log.Info("Problem0013 answer => %v", answer)

	return answer.Text(BASE)
}
