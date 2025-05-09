/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/dictionaries_and_hashmaps/sherlock_and_anagrams.md]]
 */

package hackerrank

import (
	"math/big"
	"slices"
	"strings"

	"gon.cl/algorithms/utils/log"
)

func factorialBig(n int64) *big.Int {
	if n <= 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func sherlockAndAnagrams(s string) int32 {
	candidates := make(map[string][]string)
	size := len(s)

	for i := range size {
		for j := range size - i {
			substr := s[i : size-j]

			log.Debug("i: %d, size: %d, size - j: %d | substr: %s",
				i, size, size-j, substr)

			anagramCandidateSlice := strings.Split(substr, "")
			slices.Sort(anagramCandidateSlice)
			anagramCandidate := strings.Join(anagramCandidateSlice, "")

			_, ok := candidates[anagramCandidate]
			if ok {
				// Key exists
				candidates[anagramCandidate] = append(candidates[anagramCandidate], substr)
			} else {
				// Key does not exist
				candidates[anagramCandidate] = []string{substr}
			}
		}
	}

	var total int32 = 0
	var qCandidates int32 = 0

	// # Final Anagram list
	for word := range candidates {
		quantityOfAnagrams := int32(len(candidates[word]))
		k := int32(2)

		if quantityOfAnagrams <= 1 {
			delete(candidates, word)
		} else {
			// # Binomial coefficient: https://en.wikipedia.org/wiki/Binomial_coefficient
			qCandidates += quantityOfAnagrams

			result := big.NewInt(1)

			count := result.Div(
				factorialBig(int64(quantityOfAnagrams)),
				result.Mul(
					factorialBig(int64(k)),
					factorialBig(int64(quantityOfAnagrams-k))))

			total += int32(count.Int64())

			log.Debug("Partial anagrams of %s: %d", word, count)
		}

		log.Debug(
			"sherlock_and_anagrams(%s) Filtered # candidates: %d", s, qCandidates)
		log.Debug("sherlock_and_anagrams(%s) # anagrams: %d", s, total)
	}
	return total
}

func SherlockAndAnagrams(s string) int32 {
	return sherlockAndAnagrams((s))
}
