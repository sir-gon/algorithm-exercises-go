package helpers

import "errors"

const __ERR_EMPTY_LIST__ = "list is empty"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMinOfMany(list []int) (int, error) {
	if len(list) <= 0 {
		return 0, errors.New(__ERR_EMPTY_LIST__)
	}

	minimum := list[0]

	for i := 1; i < len(list); i += 1 {
		minimum = IntMin(list[i], minimum)
	}

	return minimum, nil
}

func IntMaxOfMany(list []int) (int, error) {
	if len(list) <= 0 {
		return 0, errors.New(__ERR_EMPTY_LIST__)
	}

	maximum := list[0]

	for i := 1; i < len(list); i += 1 {
		maximum = IntMax(list[i], maximum)
	}

	return maximum, nil
}
