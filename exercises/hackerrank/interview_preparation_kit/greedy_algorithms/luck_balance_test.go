package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type LuckBalanceTestCase struct {
	Contests [][]int32 `json:"contests"`
	K        int32     `json:"k"`
	Expected int32     `json:"expected"`
}

var LuckBalanceTestCases []LuckBalanceTestCase

// You can use testing.T, if you want to test the code without benchmarking
func LuckBalanceSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/luck_balance.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &LuckBalanceTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestLuckBalance(t *testing.T) {

	LuckBalanceSetupSuite(t)

	for _, tt := range LuckBalanceTestCases {
		testname := fmt.Sprintf("luckBalance(%v, %v) => %v \n", tt.K, tt.Contests, tt.Expected)
		t.Run(testname, func(t *testing.T) {
			ans := luckBalance(tt.K, tt.Contests)
			assert.Equal(t, tt.Expected, ans)
		})
	}
}
