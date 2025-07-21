package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type FlipingBitsTest struct {
	Input    int64 `json:"input"`
	Expected int64 `json:"answer"`
}

type FlipingBitsTests struct {
	Title string            `json:"title"`
	Tests []FlipingBitsTest `json:"tests"`
}

var FlipingBitsTestCases []FlipingBitsTests

// You can use testing.T, if you want to test the code without benchmarking
func FlipingBitsSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/flipping_bits.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &FlipingBitsTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestFlipingBits(t *testing.T) {

	FlipingBitsSetupSuite(t)

	for _, tt := range FlipingBitsTestCases {
		for _, testCase := range tt.Tests {
			testname := fmt.Sprintf("FlipingBits(%d) => %d \n", testCase.Input, testCase.Expected)

			t.Run(testname, func(t *testing.T) {
				ans := FlippingBits(testCase.Input)
				assert.Equal(t, testCase.Expected, ans)
			})
		}
	}
}
