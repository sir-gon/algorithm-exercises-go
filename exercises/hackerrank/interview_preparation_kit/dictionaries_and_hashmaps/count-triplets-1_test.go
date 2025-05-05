package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type CountTripletsTestCase struct {
	Input    []int64 `json:"Input"`
	Ratio    int64   `json:"r"`
	Expected int64   `json:"expected"`
}

var CountTripletsTestCases []CountTripletsTestCase
var CountTripletsBigTestCases []CountTripletsTestCase

// You can use testing.T, if you want to test the code without benchmarking
func CountTripletsSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/count-triplets-1.small.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &CountTripletsTestCases)
	if err != nil {
		t.Log(err)
	}
}

func CountTripletsSetupSuiteBigCases(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/count-triplets-1.big.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &CountTripletsBigTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestCountTriplets(t *testing.T) {

	CountTripletsSetupSuite(t)

	for _, tt := range CountTripletsTestCases {
		testname := fmt.Sprintf("CountTriplets(%v, %v) => %v \n", tt.Input, tt.Ratio, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := CountTriplets(tt.Input, tt.Ratio)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}

func TestCountTripletsBigCases(t *testing.T) {

	CountTripletsSetupSuiteBigCases(t)

	for _, tt := range CountTripletsBigTestCases {
		testname := fmt.Sprintf("CountTriplets(%v, %v) => %v \n", tt.Input, tt.Ratio, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := CountTriplets(tt.Input, tt.Ratio)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}

func TestCountTripletsBruteForce(t *testing.T) {

	CountTripletsSetupSuiteBigCases(t)

	for _, tt := range CountTripletsTestCases {
		testname := fmt.Sprintf("CountTripletsBruteForce(%v, %v) => %v \n", tt.Input, tt.Ratio, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := CountTripletsBruteForce(tt.Input, tt.Ratio)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
