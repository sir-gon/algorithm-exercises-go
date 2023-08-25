/**
 * @link Problem definition [[docs/hackerrank/warmup/plusMinus.md]]
 */

package hackerrank

import (
	"fmt"
	"strings"

	utils "gon.cl/algorithm-exercises/src/utils"
)

func PlusMinus(arr []int) string {

	positives := 0
	negatives := 0
	zeros := 0

	for i := 0; i < len(arr); i += 1 {
		var num = arr[i]

		if num > 0 {
			positives += 1
		} else if num < 0 {
			negatives += 1
		} else if num == 0 {
			zeros += 1
		}
	}

	result := []string{}

	result = append(result, fmt.Sprintf("%.6f", float64(positives)/float64(len(arr))))
	result = append(result, fmt.Sprintf("%.6f", float64(negatives)/float64(len(arr))))
	result = append(result, fmt.Sprintf("%.6f", float64(zeros)/float64(len(arr))))

	answer := strings.Join(result, "\n")

	utils.Info("PlusMinus answer => %s", answer)

	return answer
}
