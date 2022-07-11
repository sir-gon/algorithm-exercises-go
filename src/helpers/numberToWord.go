package helpers

import (
	"fmt"
	"math/big"

	log "gon.cl/projecteuler.net/src/lib"
)

const _CENTS_ = "hundred"
const _MILLS_ = "thousand"

var dictionary = map[string]string{
	"1":    "one",
	"2":    "two",
	"3":    "three",
	"4":    "four",
	"5":    "five",
	"6":    "six",
	"7":    "seven",
	"8":    "eight",
	"9":    "nine",
	"10":   "ten",
	"11":   "eleven",
	"12":   "twelve",
	"13":   "thirteen",
	"14":   "fourteen",
	"15":   "fifteen",
	"16":   "sixteen",
	"17":   "seventeen",
	"18":   "eighteen",
	"19":   "nineteen",
	"20":   "twenty",
	"30":   "thirty",
	"40":   "forty",
	"50":   "fifty",
	"60":   "sixty",
	"70":   "seventy",
	"80":   "eighty",
	"90":   "ninety",
	"100":  "hundred",
	"1000": "thousand",
}

func NumberToWord(_value big.Int) string {
	number := _value.Text(__NUMERIC_BASE__)
	digits := len(number)

	var bottomLimit, upperLimit bool

	/// 1 to 19
	bottomLimit = _value.Cmp(big.NewInt(0)) == 1
	upperLimit = _value.Cmp(big.NewInt(19)) <= 0

	if digits <= 2 && bottomLimit && upperLimit {
		return dictionary[number]
	}

	/// 20 to 99
	bottomLimit = _value.Cmp(big.NewInt(20)) == 1
	upperLimit = _value.Cmp(big.NewInt(99)) <= 0

	if digits <= 2 && bottomLimit && upperLimit {

		dec := new(big.Int)
		div := big.NewInt(10)
		unit := new(big.Int)

		dec, unit = dec.DivMod(&_value, div, unit)
		dec = dec.Mul(dec, big.NewInt((10)))

		log.Info("dec => %s", dec.Text(__NUMERIC_BASE__))
		log.Info("div => %s", div.Text(__NUMERIC_BASE__))
		log.Info("unit => %s", unit.Text(__NUMERIC_BASE__))

		unitIsZero := unit.Cmp(big.NewInt(0))
		if unitIsZero == 0 {
			return dictionary[number]
		}

		return fmt.Sprintf("%s-%s", dictionary[dec.Text(__NUMERIC_BASE__)], dictionary[unit.Text(__NUMERIC_BASE__)])

	}

	return ""
}
