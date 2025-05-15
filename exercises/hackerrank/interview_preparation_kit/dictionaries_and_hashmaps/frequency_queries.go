/*
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/dictionaries_and_hashmaps/frequency-queries.md]]
 */

package hackerrank

const __INITIAL__ = 0

const __INSERT__ = 1
const __DELETE__ = 2
const __SELECT__ = 3

const __NOT_FOUND__ = 0
const __FOUND__ = 1

func freqQuery(queries [][]int32) []int32 {
	dataMap := make(map[int32]int32)
	freq := make(map[int32]int32)
	var result []int32

	for _, query := range queries {
		switch query[0] {
		case __INSERT__:
			if dataMap[query[1]] == __INITIAL__ {
				dataMap[query[1]] = 1
			} else {
				freq[dataMap[query[1]]]--
				dataMap[query[1]]++
			}
			freq[dataMap[query[1]]]++
		case __DELETE__:
			if dataMap[query[1]] > __INITIAL__ {
				freq[dataMap[query[1]]]--
				dataMap[query[1]]--
				freq[dataMap[query[1]]]++
			}
		case __SELECT__:
			if freq[query[1]] > 0 {
				result = append(result, __FOUND__)
			} else {
				result = append(result, __NOT_FOUND__)
			}
		}
	}

	return result
}

func FreqQuery(queries [][]int32) []int32 {
	return freqQuery(queries)
}
