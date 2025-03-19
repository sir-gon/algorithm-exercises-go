package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type CrushTestCase struct {
	N        int32     `json:"n"`
	Queries  [][]int32 `json:"queries"`
	Expected int64     `json:"expected"`
}

var CrushTestCases []CrushTestCase

// You can use testing.T, if you want to test the code without benchmarking
func CrushSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/crush.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &CrushTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestCrush(t *testing.T) {

	CrushSetupSuite(t)

	for _, tt := range CrushTestCases {
		testname := fmt.Sprintf("ArrayManipulation(%d, %v) => %d \n", tt.N, tt.Queries, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			ans := ArrayManipulation(tt.N, tt.Queries)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
