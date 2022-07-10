package helpers

import (
	"testing"
)

func TestBigSumManyBasic(t *testing.T) {

	input := []string{"1", "2", "3"}
	expected := "6"

	ans := BigSumMany(input)

	if ans != expected {

		t.Errorf("BigSumMany(%v) = %v; want %s", input, ans, expected)
	}
}

func TestBigSumManyError(t *testing.T) {

	input := []string{"1", "2", "a"}
	expected := "0"

	ans := BigSumMany(input)

	if ans != expected {

		t.Errorf("BigSumMany(%v) = %v; want %s", input, ans, expected)
	}
}
