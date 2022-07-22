package lib

type BNode struct {
	value       int
	left, right *BNode
}

func (node *BNode) setValue(_value int) BNode {
	node.value = _value
	return *node
}

func (node BNode) getValue() int {
	return node.value
}

func (node *BNode) setLeft(_left BNode) BNode {
	node.left = &_left
	return *node
}

func (node BNode) getLeft() *BNode {
	return node.left
}

func (node *BNode) setRight(_right BNode) BNode {
	node.right = &_right
	return *node
}

func (node BNode) getRight() *BNode {
	return node.right
}

func (node BNode) isLeaf() bool {
	return node.left == nil && node.right == nil
}
