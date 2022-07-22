package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {

	expectedNode := BNode{5, nil, nil}
	expectedLeafs := []int{5}

	testname := fmt.Sprintf("Tree => root: %v | leafs: %v \n", expectedNode, expectedLeafs)
	t.Run(testname, func(t *testing.T) {

		tree := Tree{
			BNode{5, nil, nil},
			[]int{5},
		}

		assert.Equal(t, expectedLeafs, tree.GetLeafs())
		assert.Equal(t, expectedNode, tree.GetRoot())
	})
}
