package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type AngryChildrenTestCase struct {
	K        int32   `json:"k"`
	Arr      []int32 `json:"arr"`
	Expected int32   `json:"expected"`
}

var AngryChildrenTestCases []AngryChildrenTestCase

// You can use testing.T, if you want to test the code without benchmarking
func AngryChildrenSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/angry_children.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &AngryChildrenTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestAngryChildren(t *testing.T) {

	AngryChildrenSetupSuite(t)

	for _, tt := range AngryChildrenTestCases {
		testname := fmt.Sprintf("MaxMin(%d, %v) => %v \n", tt.K, tt.Arr, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := MaxMin(tt.K, tt.Arr)
			assert.Equal(t, tt.Expected, ans)
		})
	}
}
