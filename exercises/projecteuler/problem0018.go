/**
 * @link Problem definition [[docs/projecteuler/problem0018.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/exercises/projecteuler/lib"
	utils "gon.cl/algorithms/utils"
)

func Problem0018(tree lib.Tree) int {

	var answer int

	answer, err := helpers.IntMaxOfMany(tree.GetLeafs())

	utils.Info("Problem0018 answer => leafs: %d", len(tree.GetLeafs()))
	utils.Info("Problem0018 answer => Max leaf: %d | err: %v", answer, err)

	return answer
}
