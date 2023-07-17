/**
 * Names scores
 *
 * https://projecteuler.net/problem=22
 *
 * Using listOfNames.txt https://projecteuler.net/project/resources/p022_listOfNames.txt
 * (right click and 'Save Link/Target As...'),
 * a 46K text file containing over five-thousand first names,
 * begin by sorting it into alphabetical order.
 * Then working out the alphabetical value for
 * each name, multiply this value by its alphabetical position
 * in the list to obtain a name score.
 *
 * For example, when the list is sorted into
 * alphabetical order, COLIN, which
 * is worth 3 + 15 + 12 + 9 + 14 = 53, is the
 * 938th name in the list. So, COLIN would obtain
 * a score of 938 × 53 = 49714.
 *
 * What is the total of all the name scores in the file?
 */

// ////////////////////////////////////////////////////////////////////////////
// See:
//    - src/data/p022_listOfNames.txt
//    - src/data/p022_listOfNames.json
// ////////////////////////////////////////////////////////////////////////////

package projecteuler

import (
	"sort"

	"gon.cl/projecteuler.net/src/projecteuler/helpers"
	"gon.cl/projecteuler.net/src/projecteuler/utils"
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
