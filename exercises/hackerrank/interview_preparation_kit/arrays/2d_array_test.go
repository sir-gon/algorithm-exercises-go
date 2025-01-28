package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type TwoDArrayTestCase struct {
	Input    [][]int32 `json:"input"`
	Expected int32     `json:"expected"`
}

var TwoDArrayTestCases []TwoDArrayTestCase

// You can use testing.T, if you want to test the code without benchmarking
func twoDArraySetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/2d_array.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &TwoDArrayTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestTwoDArray(t *testing.T) {

	twoDArraySetupSuite(t)

	for _, tt := range TwoDArrayTestCases {
		testname := fmt.Sprintf("hourglassSum(%v) => %v \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			ans := HourglassSum(tt.Input)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
