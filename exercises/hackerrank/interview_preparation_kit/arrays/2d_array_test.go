package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type TestCase struct {
	Input    [][]int32 `json:"input"`
	Expected int32     `json:"expected"`
}

var testCases []TestCase

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/2d_array.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &testCases)
	if err != nil {
		t.Log(err)
	}
}

func Test2DArray(t *testing.T) {

	setupSuite(t)

	for _, tt := range testCases {
		testname := fmt.Sprintf("hourglassSum(%v) => %v \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			ans := hourglassSum(tt.Input)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
