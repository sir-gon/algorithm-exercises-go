/**
 * @link Problem definition [[docs/projecteuler/problem0008.md]]
 */

package projecteuler

import (
	"strconv"
	"strings"

	"gon.cl/algorithms/exercises/projecteuler/helpers"
	utils "gon.cl/algorithms/utils"
)

const __BASE__ = 10
const __BIT_SIZE__ = 32

func Problem0008(numberInput string) int {

	var greatest = 0

	var digitsSlice []int32
	var bigNumSlice = strings.Split(numberInput, "")

	for i := 0; i < len(bigNumSlice); i += 1 {
		v, _ := strconv.ParseInt(string(bigNumSlice[i]), __BASE__, __BIT_SIZE__)
		digitsSlice = append(digitsSlice, int32(v))
	}

	const interval = 13
	const bottom = 0
	var top = len(digitsSlice)

	for i := bottom; i <= top-interval; i += 1 {
		var digitsSet []int

		for j := 0; j < interval; j++ {
			digitsSet = append(digitsSet, int(digitsSlice[i+j]))
		}

		var currentProduct = helpers.Product(digitsSet)
		utils.Debug("Product beetwen %d and %d <%d> is: %d", i, i+interval, digitsSet, currentProduct)

		if currentProduct > greatest {
			greatest = currentProduct
		}
	}

	utils.Info("Problem0008 => The the greatest product of %d consecutive digits is: %d", interval, greatest)

	return greatest
}
