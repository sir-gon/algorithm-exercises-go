/**
 * Longest Collatz sequence
 *
 * https://projecteuler.net/problem=14
 *
 *
 * The following iterative sequence is defined for the set of positive integers:
 *
 * n → n/2 (n is even)
 * n → 3n + 1 (n is odd)
 *
 * Using the rule above and starting with 13, we generate the following sequence:
 *
 * 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
 * It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
 * Although it has not been proved yet (Collatz Problem), it is thought that all
 * starting numbers finish at 1.
 *
 * Which starting number, under one million, produces the longest chain?
 *
 * NOTE: Once the chain starts the terms are allowed to go above one million.
 */

package projecteuler

import (
	"gon.cl/projecteuler.net/src/helpers"
	utils "gon.cl/projecteuler.net/src/utils"
)

func Problem0014(bottom int, top int) int {
	var answer int
	var maxSequence []int

	for i := bottom; i < top; i += 1 {

		var sequence []int = helpers.CollatzSequence(i)
		utils.Info("sequence of %d: %v", i, sequence)

		if len(sequence) > len(maxSequence) {
			maxSequence = sequence
		}
	}

	utils.Info("Large sequence found: ${maxSequence} has ${maxSequence.length} elements")

	// return firt element
	answer = maxSequence[0]

	utils.Info("Problem0014 answer => %d", answer)
	return answer
}
