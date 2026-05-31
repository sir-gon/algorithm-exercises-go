package hackerrank

import (
	"sort"
)

func RemainderSorting(strArr []string) []string {
	sortingmap := make(map[int][]string)
	keys := []int{}

	for i := range strArr {
		word := strArr[i]
		score := len(strArr[i]) % 3

		_, ok := sortingmap[score]
		if ok {
			sortingmap[score] = append(sortingmap[score], word)
		} else {
			sortingmap[score] = []string{word}
			keys = append(keys, score)
		}

		sort.Strings(sortingmap[score])
	}

	sort.Ints(keys)

	result := []string{}

	for j := 0; j < len(keys); j += 1 {
		element := sortingmap[j]

		result = append(result, element...)
	}

	return result
}
