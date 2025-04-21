/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/dictionaries_and_hashmaps/ctci-ransom-note.md]]
 */

package hackerrank

import "fmt"

const __YES__ = "Yes"
const __NO__ = "No"

func checkMagazineCompute(magazine []string, note []string) bool {
	dictionary := make(map[string]int)

	for _, word := range magazine {
		dictionary[word]++
	}

	for _, word := range note {
		if _, ok := dictionary[word]; ok && dictionary[word] > 0 {
			dictionary[word] -= 1
		} else {
			return false
		}
	}

	return true
}

func checkMagazineText(magazine []string, note []string) string {
	if checkMagazineCompute(magazine, note) {
		return __YES__
	}

	return __NO__
}

func checkMagazine(magazine []string, note []string) {
	fmt.Println(checkMagazineText(magazine, note))
}

func CheckMagazineText(magazine []string, note []string) string {
	return checkMagazineText(magazine, note)
}

func CheckMagazine(magazine []string, note []string) {
	checkMagazine(magazine, note)
}
