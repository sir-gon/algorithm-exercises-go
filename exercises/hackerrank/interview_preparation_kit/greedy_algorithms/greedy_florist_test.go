package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type GreedyFloristTestCase struct {
	C        []int32 `json:"c"`
	K        int32   `json:"k"`
	Expected int32   `json:"expected"`
}

var GreedyFloristTestCases []GreedyFloristTestCase

// You can use testing.T, if you want to test the code without benchmarking
func GreedyFloristSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/greedy_florist.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &GreedyFloristTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestGreedyFlorist(t *testing.T) {

	GreedyFloristSetupSuite(t)

	for _, tt := range GreedyFloristTestCases {
		testname := fmt.Sprintf("getMinimumCost(%v, %v) => %v \n", tt.K, tt.C, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := GetMinimumCost(tt.K, tt.C)
			assert.Equal(t, tt.Expected, ans)
		})
	}
}
