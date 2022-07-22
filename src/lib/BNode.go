package lib

type BNode struct {
	value       int
	left, right *BNode
}

func (node *BNode) setValue(_value int) {
	node.value = _value
}

func (node BNode) getValue() int {
	return node.value
}

func (node *BNode) setLeft(_left BNode) {
	node.left = &_left
}

func (node BNode) getLeft() *BNode {
	return node.left
}

func (node *BNode) setRight(_right BNode) {
	node.right = &_right
}

func (node BNode) getRight() *BNode {
	return node.right
}
