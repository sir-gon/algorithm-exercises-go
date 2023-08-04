package projecteuler

import (
	"gon.cl/algorithm-exercises/src/pkg/projecteuler/helpers"
	utils "gon.cl/algorithm-exercises/src/utils"
)

func Problem0014(bottom int, top int) int {
	var answer int
	var maxSequence []int

	for i := bottom; i < top; i += 1 {

		var sequence []int = helpers.CollatzSequence(i)
		utils.Debug("sequence of %d: %v", i, sequence)

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
