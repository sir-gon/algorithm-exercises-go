/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/dictionaries_and_hashmaps/count_triplets_1.md]]
 */

package hackerrank

func countTripletsBruteForce(arr []int64, r int64) int64 {
	size := len(arr)
	counter := int64(0)
	for i := range size - 2 {
		for j := i + 1; j < size-1; j++ {
			for k := j + 1; k < size; k++ {
				if r*arr[i] == arr[j] && r*arr[j] == arr[k] {
					counter++
				}
			}
		}
	}
	return counter
}

func countTriplets(arr []int64, r int64) int64 {
	aCounter := make(map[int64]int64)
	bCounter := make(map[int64]int64)
	triplets := int64(0)

	for _, item := range arr {
		if _, ok := aCounter[item]; ok && aCounter[item] > 0 {
			aCounter[item] += 1
		} else {
			aCounter[item] = 1
		}
	}

	for _, item := range arr {
		j := item / r
		k := item * r

		aCounter[item]--
		if bCounter[j] > 0 && aCounter[k] > 0 && item%r == 0 {
			triplets += bCounter[j] * aCounter[k]
		}
		bCounter[item]++
	}

	return triplets
}

func CountTriplets(arr []int64, r int64) int64 {
	return countTriplets(arr, r)
}

func CountTripletsBruteForce(arr []int64, r int64) int64 {
	return countTripletsBruteForce(arr, r)
}
