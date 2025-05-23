/**
 * @link Problem definition [[docs/hackerrank/warmup/staircase.md]]
 */

package hackerrank

import (
	"strings"

	"gon.cl/algorithms/utils/log"
)

func Staircase(n int) string {

	result := []string{}

	for i := 1; i <= n; i++ {
		line := ""

		for j := 1; j <= n; j++ {
			if j <= n-i {
				line = line + " "
			} else {
				line = line + "#"
			}
		}

		result = append(result, line)
	}

	log.Info("Staircase answer => %v", result)

	return strings.Join(result, "\n")
}
