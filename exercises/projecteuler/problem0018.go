/**
 * @link Problem definition [[docs/projecteuler/problem0018.md]]
 */

package projecteuler

import (
	"gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/exercises/projecteuler/lib"
	"gon.cl/algorithms/utils/log"
)

func Problem0018(tree lib.Tree) int {

	var answer int

	answer, err := helpers.IntMaxOfMany(tree.GetLeafs())

	log.Info("Problem0018 answer => leafs: %d", len(tree.GetLeafs()))
	log.Info("Problem0018 answer => Max leaf: %d | err: %v", answer, err)

	return answer
}
