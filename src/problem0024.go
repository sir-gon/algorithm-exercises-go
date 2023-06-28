/**
 * Lexicographic Permutations
 *
 * https://projecteuler.net/problem=24
 *
 * A permutation is an ordered arrangement of objects.
 * For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
 * If all of the permutations are listed numerically or alphabetically,
 * we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
 *
 * 012   021   102   120   201   210
 *
 * What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
 */

package projecteuler

import (
	utils "gon.cl/projecteuler.net/src/utils"
)

func lexicographicPermutationFind(elements []string, permutationToFind int) []string {
	// "inner global variables" to catch the results shared across recursive branch calls.
	var lastBranchCollector []string = []string{}
	var currentCycle = 0

	// Initial values
	var initBranchCollector []string = []string{}

	// Recursive way to compute permutations
	var computePermutations func(stopAtCycle int, inputElements []string, branchCollector []string)

	computePermutations = func(stopAtCycle int, inputElements []string, branchCollector []string) {
		if currentCycle >= stopAtCycle {
			return
		}

		for i := 0; i < len(inputElements); i++ {
			var rootElement = inputElements[i]
			var restOfElements []string = []string{}

			utils.Debug("root element: %d -> %s", i, rootElement)

			for j := 0; j < len(inputElements); j++ {
				if i != j {
					restOfElements = append(restOfElements, inputElements[j])
				}
			}

			newBranchCollector := make([]string, len(branchCollector))
			copy(newBranchCollector, branchCollector)
      newBranchCollector = append(newBranchCollector, rootElement)

			// finally...
			if len(restOfElements) > 0 {
				utils.Debug("REST: %+v", restOfElements)

				computePermutations(stopAtCycle, restOfElements, newBranchCollector)
			} else {
				lastBranchCollector = newBranchCollector
				currentCycle += 1

				utils.Debug("FINISH BRANCH: %d -> %+v", currentCycle, lastBranchCollector)
			}
		}
	}

	computePermutations(permutationToFind, elements, initBranchCollector)

	return lastBranchCollector
}

func Problem0024(inputElements []string, inputPermutationToFind int) []string {

	var permutationFound = lexicographicPermutationFind(inputElements, inputPermutationToFind)

	utils.Info("Problem0024 answer => %+v", permutationFound)

	return permutationFound
}
