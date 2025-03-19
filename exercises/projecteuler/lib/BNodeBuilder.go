package lib

import "gon.cl/algorithms/utils"

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

func buildBNodeTreeWeigthRecursive(
	dataTree [][]int,
	i int,
	j int,
	rootValue BNode,
	leafCollector []int,
) (result BNode, leaves []int, err bool) {
	if 0 <= i && i < len(dataTree) && 0 <= j && j < len(dataTree[i]) {

		resultNode := BNode{dataTree[i][j] + rootValue.getValue(), nil, nil}

		if i+1 <= len(dataTree)-1 && j <= len(dataTree[i+1])-1 &&
			i+1 <= len(dataTree)-1 && j+1 <= len(dataTree[i+1])-1 {

			var left BNode
			var right BNode
			var err bool

			left, leafCollector, err = buildBNodeTreeWeigthRecursive(dataTree, i+1, j, resultNode, leafCollector)
			utils.Debug("left: %v | err: %t | leavesCarry: %v", left, err, leafCollector)

			resultNode.setLeft(left)

			right, leafCollector, err = buildBNodeTreeWeigthRecursive(dataTree, i+1, j+1, resultNode, leafCollector)
			utils.Debug("right: %v | err: %t | leavesCarry: %v", left, err, leafCollector)

			resultNode.setRight(right)

		}

		if resultNode.isLeaf() {
			leafCollector = append(leafCollector, resultNode.getValue())
		}

		return resultNode, leafCollector, false
	}
	return BNode{}, []int{}, true
}

func BuildBNodeTreeWeigth(dataTree [][]int) (result Tree, err bool) {
	leaves := []int{}
	root := BNode{0, nil, nil}

	newRoot, newLeaves, err := buildBNodeTreeWeigthRecursive(dataTree, 0, 0, root, leaves)

	return Tree{newRoot, newLeaves}, err
}
