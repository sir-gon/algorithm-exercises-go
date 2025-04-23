/**
 * @link Problem definition [[docs/hackerrank/warmup/timeConversion.md]]
 */

package hackerrank

import (
	"fmt"
	"strconv"
	"strings"

	"gon.cl/algorithms/utils/log"
)

func TimeConversion(s string) string {
	meridian := s[len(s)-2:]
	meridian = strings.ToLower(meridian)

	timeStr := s[0 : len(s)-2]
	time := strings.Split(timeStr, ":")
	hour, _ := strconv.Atoi(time[0])

	if hour >= 12 {
		hour = 0
	}

	if meridian == "pm" {
		hour += 12
	}

	time[0] = fmt.Sprintf("%02d", hour)

	conversion := strings.Join(time, ":")

	log.Info("TimeConversion answer => %s", conversion)

	return conversion
}
