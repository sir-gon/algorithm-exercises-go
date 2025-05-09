package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type MinimumAbsoluteDifferenceInAnArrayTestCase struct {
	Input    []int32 `json:"input"`
	Expected int32   `json:"expected"`
}

var MinimumAbsoluteDifferenceInAnArrayTestCases []MinimumAbsoluteDifferenceInAnArrayTestCase
var MinimumAbsoluteDifferenceInAnArrayTestCase2 []MinimumAbsoluteDifferenceInAnArrayTestCase

// You can use testing.T, if you want to test the code without benchmarking
func MinimumAbsoluteDifferenceInAnArraySetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/minimum_absolute_difference_in_an_array.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &MinimumAbsoluteDifferenceInAnArrayTestCases)
	if err != nil {
		t.Log(err)
	}
}

func MinimumAbsoluteDifferenceInAnArraySetupSuiteTestCase2(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/minimum_absolute_difference_in_an_array.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &MinimumAbsoluteDifferenceInAnArrayTestCase2)
	if err != nil {
		t.Log(err)
	}
}

func TestMinimumAbsoluteDifferenceInAnArray(t *testing.T) {

	MinimumAbsoluteDifferenceInAnArraySetupSuite(t)

	for _, tt := range MinimumAbsoluteDifferenceInAnArrayTestCases {
		testname := fmt.Sprintf("MinimumAbsoluteDifferenceInAnArray(%v) => %v \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := MinimumAbsoluteDifferenceInAnArray(tt.Input)
			assert.Equal(t, tt.Expected, ans)
		})
	}
}

func TestMinimumAbsoluteDifferenceInAnArrayBigCase(t *testing.T) {

	MinimumAbsoluteDifferenceInAnArraySetupSuiteTestCase2(t)

	for _, tt := range MinimumAbsoluteDifferenceInAnArrayTestCase2 {
		testname := fmt.Sprintf("MinimumAbsoluteDifferenceInAnArray(%v) => %v \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := MinimumAbsoluteDifferenceInAnArray(tt.Input)
			assert.Equal(t, tt.Expected, ans)
		})
	}
}
