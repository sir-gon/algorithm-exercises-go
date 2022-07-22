package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBNode(t *testing.T) {

	const expectedValue = 1

	testname := fmt.Sprintf("BNode() => %v \n", expectedValue)
	t.Run(testname, func(t *testing.T) {

		ans := BNode{1, nil, nil}

		assert.Equal(t, expectedValue, ans.getValue())
	})
}

func TestBNodeValue(t *testing.T) {

	const expectedValue = 1
	x := BNode{0, nil, nil}

	testname := fmt.Sprintf("BNode() => %v \n", expectedValue)
	t.Run(testname, func(t *testing.T) {

		x.setValue(1)

		assert.Equal(t, expectedValue, x.getValue())
	})
}

func TestBNodeLeft(t *testing.T) {

	const expectedLeftValue = 2

	y := BNode{0, nil, nil}
	left := BNode{2, nil, nil}

	testname := fmt.Sprintf("BNode() => %v \n", y)
	t.Run(testname, func(t *testing.T) {

		y.setLeft(left)

		fmt.Printf("LEFT: %v\n", left)
		fmt.Printf("y: %v\n", y)

		assert.Equal(t, expectedLeftValue, y.getLeft().getValue())

	})
}

func TestBNodeRight(t *testing.T) {

	const expectedLeftValue = 2

	y := BNode{0, nil, nil}
	right := BNode{2, nil, nil}

	testname := fmt.Sprintf("BNode() => %v \n", y)
	t.Run(testname, func(t *testing.T) {

		y.setRight(right)

		fmt.Printf("RIGHT: %v\n", right)
		fmt.Printf("y: %v\n", y)

		assert.Equal(t, expectedLeftValue, y.getRight().getValue())

	})
}
