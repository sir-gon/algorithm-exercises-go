/**
 * Largest palindrome product
 *
 * https://projecteuler.net/problem=4
 *
 * A palindromic number reads the same both ways.
 * The largest palindrome made from the product of two 2-digit
 * numbers is 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */

package projecteuler

import (
	"gon.cl/projecteuler.net/src/helpers"
	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0004(_bottom int, _top int) int {
	var i int
	var j int
	var foundi int
	var foundj int
	var foundPalindrome int

	log.Info("Initializing Problem 0004")

	i = _top
	for oki := true; oki; oki = i >= _bottom {
		j = _top
		for okj := true; okj; okj = j >= _bottom {

			if helpers.IsPalindrome(j * i) {

				log.Debug("FOUND %d x %d = %d is Palindrome", i, j, i*j)

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

	log.Info("Problem0004 Largest Palindrome => %d 𝗑 %d = %d", foundi, foundj, foundPalindrome)

	return foundPalindrome
}
