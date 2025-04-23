package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type TwoStringsTestCase struct {
	S1       string `json:"s1"`
	S2       string `json:"s2"`
	Expected string `json:"expected"`
}

var TwoStringsTestCases []TwoStringsTestCase

// You can use testing.T, if you want to test the code without benchmarking
func TwoStringsSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/two_strings.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &TwoStringsTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestTwoStrings(t *testing.T) {

	TwoStringsSetupSuite(t)

	for _, tt := range TwoStringsTestCases {
		testname := fmt.Sprintf("TwoStrings(%v, %v) => %v \n", tt.S1, tt.S2, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := TwoStrings(tt.S1, tt.S2)
			assert.Equal(t, tt.Expected, ans)
		})
	}
}
