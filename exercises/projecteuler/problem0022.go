/**
 * @link Problem definition [[docs/projecteuler/problem0012.md]]
 * @see [[src/pkg/projecteuler/data/p022_names.txt]]
 * @see [[src/pkg/projecteuler/data/problem0022_data.go]]
 */

package projecteuler

import (
	"sort"

	"gon.cl/algorithms/exercises/projecteuler/helpers"
	"gon.cl/algorithms/utils/log"
)

func Problem0022(listOfNames []string) int {

	var answer int

	sort.Strings(listOfNames)

	for i := 0; i < len(listOfNames); i++ {
		answer += (i + 1) * helpers.WordScore(listOfNames[i])
	}

	log.Info("Problem0022 answer => %d", answer)

	return answer
}

// export function problem0022(listOfNames) {
//   listOfNames.sort((a, b) => a.localeCompare(b));

//   let result = 0;
//   let i;

//   for (i = 0; i < listOfNames.length; i++) {
//     result += (i + 1) * wordScore(listOfNames[i]);
//   }

//   logger.info(`result`, result);

//   return result;
// }
