package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type TestRotLeftTestCase struct {
	Input       []int32 `json:"input"`
	D_Rotations int32   `json:"d_rotations"`
	Expected    []int32 `json:"expected"`
}

var TestRotLefttestCases []TestRotLeftTestCase

// You can use testing.T, if you want to test the code without benchmarking
func testRotLeftSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/ctci_array_left_rotation.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &TestRotLefttestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestRotLeft(t *testing.T) {

	testRotLeftSetupSuite(t)

	for _, tt := range TestRotLefttestCases {
		testname := fmt.Sprintf("RotLeft(%v) => %v \n", tt.Input, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			ans := RotLeft(tt.Input, tt.D_Rotations)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
