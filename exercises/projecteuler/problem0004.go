/**
 * @link Problem definition [[docs/projecteuler/problem0004.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	utils "gon.cl/algorithms/utils"
)

func Problem0004(bottom int, top int) int {
	var i int
	var j int
	var foundi int
	var foundj int
	var foundPalindrome int
	var cycles = 0

	utils.Info("Initializing Problem 0004")

	i = top
	for i >= bottom {
		j = i
		for j >= bottom && (foundj == 0 || j >= foundj) {

			cycles += 1
			if helpers.IsPalindrome(j * i) {

				utils.Debug("FOUND %d x %d = %d is Palindrome", i, j, i*j)

				if foundPalindrome == 0 || i*j > foundPalindrome {
					foundi = i
					foundj = j
					foundPalindrome = i * j
				}

			}

			j -= 1
		}

		i -= 1
	}

	utils.Info("Problem0004 Largest Palindrome => %d 𝗑 %d = %d in %d cycles", foundi, foundj, foundPalindrome, cycles)

	return foundPalindrome
}
