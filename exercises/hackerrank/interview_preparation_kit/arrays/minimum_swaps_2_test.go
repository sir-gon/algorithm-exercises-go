package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type MinimumTwoSwaps2TestCase struct {
	Input    []int32 `json:"input"`
	Expected int32   `json:"expected"`
}

var MinimumTwoSwaps2testCases []MinimumTwoSwaps2TestCase

// You can use testing.T, if you want to test the code without benchmarking
func testMinimumTwoSwaps2SetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/minimum_swaps_2.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &MinimumTwoSwaps2testCases)
	if err != nil {
		t.Log(err)
	}
}

func TestMinimumTwoSwaps2(t *testing.T) {

	testMinimumTwoSwaps2SetupSuite(t)

	for _, tt := range MinimumTwoSwaps2testCases {
		testname := fmt.Sprintf("MinimumSwaps(%v) => %d \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := MinimumSwaps(tt.Input)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
