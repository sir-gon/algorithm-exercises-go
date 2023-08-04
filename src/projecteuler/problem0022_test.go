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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithm-exercises/src/projecteuler/data"
)

func TestProblem0022(t *testing.T) {

	expectedSolution := 871198282

	testname := fmt.Sprintf("Problem0022(data.Problem0022_data) => %v \n", expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0022(data.Problem0022_data)
		assert.Equal(t, expectedSolution, ans)
	})
}
