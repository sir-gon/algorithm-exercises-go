/**
 * Smallest multiple
 *
 * https://projecteuler.net/problem=5
 *
 *
 *  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 *
 *  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */

package projecteuler

import (
	"math"

	"gon.cl/algorithm-exercises/src/pkg/projecteuler/helpers"
	"gon.cl/algorithm-exercises/src/utils"
)

func primeFactorListCollection(factors []int) map[int]int {
	var collection = map[int]int{}

	for _, factor := range factors {
		quantity, ok := collection[factor]

		if ok {
			collection[factor] += 1
			utils.Debug("Factor %d with quantity %d found in %v", factor, quantity, collection)
		} else {
			collection[factor] = 1
			utils.Debug("Factor %d with quantity not found in %v", factor, collection)

		}
	}

	return collection
}

func Problem0005(bottom int, top int) int {

	var minimumPrimeFactors = map[int]int{}
	var cycles = 0
	var answer = 1

	for i := bottom; i <= top; i++ {
		primeFactorList, subCycles := helpers.PrimeFactors(i)
		cycles += subCycles
		utils.Info("Prime Factors of %d list    => %v", i, primeFactorList)

		primeFactorMap := primeFactorListCollection(primeFactorList)
		cycles += len(primeFactorList)
		utils.Info("Prime Factors of %d grouped => %v", i, primeFactorMap)

		for factor, quantity := range primeFactorMap {
			cycles += 1

			elem, ok := minimumPrimeFactors[factor]

			if ok {
				minimumPrimeFactors[factor] = helpers.IntMax(quantity, elem)
			} else {
				minimumPrimeFactors[factor] = quantity
			}

		}

		utils.Info("Minimum Prime Factors of grouped until %d => %v", i, minimumPrimeFactors)
	}

	for factor, quantity := range minimumPrimeFactors {
		cycles += 1
		answer *= int(math.Pow(float64(factor), float64(quantity)))
	}

	utils.Info("Problem 0005: Minimum Prime Factors from %d to %d => %v in %d cycles", bottom, top, minimumPrimeFactors, cycles)
	utils.Info("Problem 0005: Solution: %d in %d cycles", answer, cycles)

	return answer
}