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
	answer := ""
	min := 0

	for len(choices) > 0 {
		index := 0
		combos := factorial(len(choices) - 1)
		min += combos
		for target > min {
			index += 1
			min += combos
		}
		answer += choices[index]
		copy(choices[index:], choices[index+1:])
		choices[len(choices)-1] = ""
		choices = choices[:len(choices)-1]
		min -= combos
	}

	return answer
}

func Problem0024(inputElements string, inputPermutationToFind int) string {

	var permutationFound = permute(inputElements, inputPermutationToFind)

	log.Info("Problem0024 answer => %+v", permutationFound)

	return permutationFound
}
