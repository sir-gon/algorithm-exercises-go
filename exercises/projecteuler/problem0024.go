/**
 * @link Problem definition [[docs/projecteuler/problem0024.md]]
 */

package projecteuler

import (
	"strings"

	"gon.cl/algorithms/utils/log"
)

func factorial(n int) int {
	i := n
	out := 1
	for i > 1 {
		out *= i
		i -= 1
	}
	return out
}

func permute(symbols string, target int) string {
	choices := strings.Split(symbols, "")
	var answer strings.Builder
	minimum := 0

	for len(choices) > 0 {
		index := 0
		combos := factorial(len(choices) - 1)
		minimum += combos
		for target > minimum {
			index += 1
			minimum += combos
		}
		answer.WriteString(choices[index])
		copy(choices[index:], choices[index+1:])
		choices[len(choices)-1] = ""
		choices = choices[:len(choices)-1]
		minimum -= combos
	}

	return answer.String()
}

func Problem0024(inputElements string, inputPermutationToFind int) string {

	var permutationFound = permute(inputElements, inputPermutationToFind)

	log.Info("Problem0024 answer => %+v", permutationFound)

	return permutationFound
}
