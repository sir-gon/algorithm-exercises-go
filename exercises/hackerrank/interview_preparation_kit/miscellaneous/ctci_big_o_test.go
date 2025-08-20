package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type TimeComplexityPrimalityTest struct {
	Input    int32  `json:"input"`
	Expected string `json:"answer"`
}

type TimeComplexityPrimalityTests struct {
	Title string                        `json:"title"`
	Tests []TimeComplexityPrimalityTest `json:"tests"`
}

var TimeComplexityPrimalityTestCases []TimeComplexityPrimalityTests

// You can use testing.T, if you want to test the code without benchmarking
func TimeComplexityPrimalitySetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/ctci_big_o.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &TimeComplexityPrimalityTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestTimeComplexityPrimality(t *testing.T) {

	TimeComplexityPrimalitySetupSuite(t)

	for _, tt := range TimeComplexityPrimalityTestCases {
		for _, testCase := range tt.Tests {
			testname := fmt.Sprintf("Primality(%d) => %s \n", testCase.Input, testCase.Expected)

			t.Run(testname, func(t *testing.T) {
				ans := Primality(testCase.Input)
				assert.Equal(t, testCase.Expected, ans)
			})
		}
	}
}
