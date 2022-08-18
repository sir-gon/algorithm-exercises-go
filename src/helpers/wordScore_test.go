package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordScoreTableDriven(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{input: "ABC", want: 1 + 2 + 3},
		{input: "OTTO", want: 15*2 + 20*2},
		{input: "COLIN", want: 53},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("WordScore(%s) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := WordScore(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}

func TestWordScoreUnhappyCases(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{input: "", want: 0},
		{input: "1", want: 0},
		{input: "#", want: 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("WordScore(%s) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := WordScore(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
