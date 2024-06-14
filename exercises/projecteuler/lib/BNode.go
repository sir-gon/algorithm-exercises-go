package lib

type BNode struct {
	value       int
	left, right *BNode
}

func (node *BNode) setValue(value int) BNode {
	node.value = value
	return *node
}

func (node BNode) getValue() int {
	return node.value
}

func (node *BNode) setLeft(left BNode) BNode {
	node.left = &left
	return *node
}

func (node BNode) getLeft() *BNode {
	return node.left
}

func (node *BNode) setRight(right BNode) BNode {
	node.right = &right
	return *node
}

func (node BNode) getRight() *BNode {
	return node.right
}

func (node BNode) isLeaf() bool {
	return node.left == nil && node.right == nil
}
