package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type SherlockAndAnagramsTest struct {
	Input    string `json:"input"`
	Expected int32  `json:"expected"`
}

type SherlockAndAnagramsTests struct {
	Title string                    `json:"title"`
	Tests []SherlockAndAnagramsTest `json:"tests"`
}

var SherlockAndAnagramsTestCases []SherlockAndAnagramsTests

// You can use testing.T, if you want to test the code without benchmarking
func SherlockAndAnagramsSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/sherlock-and-anagrams.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &SherlockAndAnagramsTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestSherlockAndAnagrams(t *testing.T) {

	SherlockAndAnagramsSetupSuite(t)

	for _, tt := range SherlockAndAnagramsTestCases {
		for _, testCase := range tt.Tests {
			testname := fmt.Sprintf("SherlockAndAnagrams(%s) => %d \n", testCase.Input, testCase.Expected)

			t.Run(testname, func(t *testing.T) {
				ans := SherlockAndAnagrams(testCase.Input)
				assert.Equal(t, testCase.Expected, ans)
			})
		}
	}
}
