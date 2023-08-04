package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gon.cl/algorithm-exercises/src/projecteuler/data"
)

const testBNodeBuilderLogMessage = "BNode.() => %v \n"

func TestTreeBuilding(t *testing.T) {
	// {3},
	// {7, 4}, => {7, 4}
	// {2, 4, 6} => {2, 4, 4, 6}
	// {8, 5, 5, 9, 3}, => {8, 5, 5, 9, 5, 9, 9, 3}

	expectedValue := BNode{3,
		&BNode{7,
			&BNode{2,
				&BNode{8, nil, nil},
				&BNode{5, nil, nil},
			},
			&BNode{4,
				&BNode{5, nil, nil},
				&BNode{9, nil, nil},
			}},

		&BNode{4,
			&BNode{4,
				&BNode{5, nil, nil},
				&BNode{9, nil, nil},
			},
			&BNode{6,
				&BNode{9, nil, nil},
				&BNode{3, nil, nil},
			}},
	}

	testname := fmt.Sprintf(testBNodeBuilderLogMessage, expectedValue)
	t.Run(testname, func(t *testing.T) {

		ans, err := BuildBNodeTree(data.Problem0018_small_data)

		assert.Equal(t, false, err)
		assert.Equal(t, expectedValue, ans)
	})
}

func TestTreeBorderCase(t *testing.T) {

	expectedValue := BNode{}
	expectedErr := true

	testname := fmt.Sprintf(testBNodeBuilderLogMessage, expectedValue)
	t.Run(testname, func(t *testing.T) {

		ans, err := BuildBNodeTree([][]int{})

		assert.Equal(t, expectedErr, err)
		assert.Equal(t, expectedValue, ans)
	})
}

func TestTreeBuildingWithWeigths(t *testing.T) {
	// {3},
	// {7, 4}, => {10, 7}
	// {2, 4, 6} => {12, 14, 11, 13}
	// {8, 5, 5, 9, 3}, => {20, 17, 19, 23, 16, 20, 22, 16}

	expectedValue := BNode{3,
		&BNode{10,
			&BNode{12,
				&BNode{20, nil, nil},
				&BNode{17, nil, nil},
			},
			&BNode{14,
				&BNode{19, nil, nil},
				&BNode{23, nil, nil},
			}},

		&BNode{7,
			&BNode{11,
				&BNode{16, nil, nil},
				&BNode{20, nil, nil},
			},
			&BNode{13,
				&BNode{22, nil, nil},
				&BNode{16, nil, nil},
			}},
	}

	expectedLeafs := []int{20, 17, 19, 23, 16, 20, 22, 16}
	expectedError := false

	testname := fmt.Sprintf(testBNodeBuilderLogMessage, expectedValue)
	t.Run(testname, func(t *testing.T) {

		ans, err := BuildBNodeTreeWeigth(data.Problem0018_small_data)

		assert.Equal(t, expectedError, err)
		assert.Equal(t, expectedValue, ans.root)
		assert.Equal(t, expectedLeafs, ans.leafs)
	})
}

func TestTreeBuildingWithWeigthsBorderCase(t *testing.T) {

	inputDataTree := [][]int{}

	expectedValue := BNode{}
	expectedLeafs := []int{}
	expectedError := true

	testname := fmt.Sprintf(testBNodeBuilderLogMessage, expectedValue)
	t.Run(testname, func(t *testing.T) {

		ans, err := BuildBNodeTreeWeigth(inputDataTree)

		assert.Equal(t, expectedError, err)
		assert.Equal(t, expectedValue, ans.root)
		assert.Equal(t, expectedLeafs, ans.leafs)
	})
}
