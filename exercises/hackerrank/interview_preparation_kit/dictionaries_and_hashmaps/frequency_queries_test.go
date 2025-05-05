package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type FrequencyQueriesTestCase struct {
	Input    [][]int32 `json:"input"`
	Expected []int32   `json:"expected"`
}

var FrequencyQueriesTestCases []FrequencyQueriesTestCase

// You can use testing.T, if you want to test the code without benchmarking
func FrequencyQueriesSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/frequency_queries.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &FrequencyQueriesTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestFrequencyQueries(t *testing.T) {

	FrequencyQueriesSetupSuite(t)

	for _, tt := range FrequencyQueriesTestCases {
		testname := fmt.Sprintf("FreqQuery(%v) => %v \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			ans := FreqQuery(tt.Input)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
