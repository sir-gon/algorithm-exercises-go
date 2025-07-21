/**
 * @link Problem definition [[docs/hackerrank/interview_preparation_kit/miscellaneous/flipping-bits.md]]
 */

package hackerrank

func flippingBits(n int64) int64 {
	return ^n & 0xFFFFFFFF // Flip all bits and mask to 32 bits
}

func FlippingBits(n int64) int64 {
	return flippingBits(n)
}
