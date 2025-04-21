package hackerrank

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/utils"
)

type RansomNoteTestCase struct {
	Magazine []string `json:"magazine"`
	Note     []string `json:"note"`
	Expected string   `json:"expected"`
}

var RansomNoteTestCases []RansomNoteTestCase

// You can use testing.T, if you want to test the code without benchmarking
func RansomNoteSetupSuite(t testing.TB) {
	wd, _ := os.Getwd()
	filepath := wd + "/ctci_ransom_note.testcases.json"
	t.Log("Setup test cases from JSON: ", filepath)

	var _, err = utils.LoadJSON(filepath, &RansomNoteTestCases)
	if err != nil {
		t.Log(err)
	}
}

func TestRansomNote(t *testing.T) {

	RansomNoteSetupSuite(t)

	for _, tt := range RansomNoteTestCases {
		testname := fmt.Sprintf("CheckMagazine(%v, %v) => %v \n", tt.Magazine, tt.Note, tt.Expected)
		t.Run(testname, func(t *testing.T) {

			CheckMagazine(tt.Magazine, tt.Note)
			ans := CheckMagazineText(tt.Magazine, tt.Note)
			assert.Equal(t, tt.Expected, ans)
		})

	}
}
