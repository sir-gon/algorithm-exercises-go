// ////////////////////////////////////////////////////////////////////////////
// See:
//    - src/data/p022_listOfNames.txt
//    - src/data/p022_listOfNames.json
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"sort"

	"gon.cl/algorithm-exercises/src/pkg/projecteuler/helpers"
	"gon.cl/algorithm-exercises/src/utils"
)

func Problem0022(listOfNames []string) int {

	var answer int

	sort.Strings(listOfNames)

	for i := 0; i < len(listOfNames); i++ {
		answer += (i + 1) * helpers.WordScore(listOfNames[i])
	}

	utils.Info("Problem0022 answer => %d", answer)

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
