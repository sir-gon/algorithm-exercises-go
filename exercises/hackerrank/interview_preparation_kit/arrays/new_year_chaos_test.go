package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type TestNewYearChaosTestCase struct {
	Input    []int32 `json:"input"`
	Expected string  `json:"expected"`
}

var TestNewYearChaostestCases []TestNewYearChaosTestCase

// You can use testing.T, if you want to test the code without benchmarking
func testNewYearChaosSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/new_year_chaos.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &TestNewYearChaostestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestNewYearChaos(t *testing.T) {

	testNewYearChaosSetupSuite(t)

	for _, tt := range TestNewYearChaostestCases {
		testname := fmt.Sprintf("MinimumBribes(%v) => %v \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			MinimumBribes(tt.Input)

			ans := MinimumBribesText(tt.Input)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
