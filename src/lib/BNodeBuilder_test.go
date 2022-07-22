package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	data "gon.cl/projecteuler.net/src/data"
)

func TestTreeBuilding(t *testing.T) {
	// {3},
	// {7, 4},
	// {2, 4, 6},
	// {8, 5, 9, 3},

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

	testname := fmt.Sprintf("BNode.() => %v \n", expectedValue)
	t.Run(testname, func(t *testing.T) {

		ans, err := BuildBNodeTree(data.Problem0018_small_data)

		assert.Equal(t, false, err)
		assert.Equal(t, expectedValue, ans)
	})
}

func TestTreeBorderCase(t *testing.T) {

	expectedValue := BNode{}
	expectedErr := true

	testname := fmt.Sprintf("BNode.() => %v \n", expectedValue)
	t.Run(testname, func(t *testing.T) {

		ans, err := BuildBNodeTree([][]int{})

		assert.Equal(t, expectedErr, err)
		assert.Equal(t, expectedValue, ans)
	})
}
