package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type CrushBruteForceTestCase struct {
	N        int32     `json:"n"`
	Queries  [][]int32 `json:"queries"`
	Expected int64     `json:"expected"`
}

var CrushBruteForceTestCases []CrushBruteForceTestCase

// You can use testing.T, if you want to test the code without benchmarking
func CrushBruteForceSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/crush.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &CrushBruteForceTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestCrushBruteForce(t *testing.T) {

	CrushBruteForceSetupSuite(t)

	for _, tt := range CrushBruteForceTestCases {
		testname := fmt.Sprintf("ArrayManipulationBruteForce(%d, %v) => %d \n", tt.N, tt.Queries, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			ans := ArrayManipulationBruteForce(tt.N, tt.Queries)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
