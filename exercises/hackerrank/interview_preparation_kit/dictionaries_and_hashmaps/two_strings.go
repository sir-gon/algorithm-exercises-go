/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/dictionaries_and_hashmaps/two-strings.md]]
 * @see Solution Notes [[docs/hackerrank/interview_preparation_kit/dictionaries_and_hashmaps/two-strings-solution-notes.md]]
 */

package hackerrank

const __TWO_STRINGS_YES__ = "Yes"
const __TWO_STRINGS_NO__ = "No"

func twoStringsCompute(s1, s2 string) bool {
	for _, letter := range s1 {
		for _, n := range s2 {
			if letter == n {
				return true
			}
		}
	}

	return false
}

func twoStrings(s1, s2 string) string {
	if twoStringsCompute(s1, s2) {
		return __TWO_STRINGS_YES__
	}

	return __TWO_STRINGS_NO__
}

func TwoStrings(s1, s2 string) string {
	return twoStrings(s1, s2)
}
