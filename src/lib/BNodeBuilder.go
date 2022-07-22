package lib

import "gon.cl/projecteuler.net/src/utils"

func buildBNodeTreeRecursive(
	dataTree [][]int,
	i int,
	j int,
) (result BNode, err bool) {

	if 0 <= i && i < len(dataTree) && 0 <= j && j < len(dataTree[i]) {

		resultNode := BNode{dataTree[i][j], nil, nil}

		if i+1 <= len(dataTree)-1 && j <= len(dataTree[i+1])-1 &&
			i+1 <= len(dataTree)-1 && j+1 <= len(dataTree[i+1])-1 {

			left, err := buildBNodeTreeRecursive(dataTree, i+1, j)
			utils.Debug("left: %v | err: %t", left, err)

			resultNode.setLeft(left)

			right, err := buildBNodeTreeRecursive(dataTree, i+1, j+1)
			utils.Debug("right: %v | err: %t", left, err)

			resultNode.setRight(right)

		}

		return resultNode, false
	}
	return BNode{}, true
}

func BuildBNodeTree(dataTree [][]int) (result BNode, err bool) {
	return buildBNodeTreeRecursive(dataTree, 0, 0)
}
