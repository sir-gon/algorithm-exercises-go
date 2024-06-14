package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConversionExample(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{input: "12:01:00PM", want: "12:01:00"},
		{input: "12:01:00AM", want: "00:01:00"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("TimeConversion(%s) => %s", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := TimeConversion(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}

func TestTimeConversionTestCases(t *testing.T) {
	var tests = []struct {
		title string
		input string
		want  string
	}{
		{title: "Case 0", input: "07:05:45PM", want: "19:05:45"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("TimeConversion(%s) %s => %s", tt.input, tt.title, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := TimeConversion(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
