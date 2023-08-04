package lib

type Tree struct {
	root  BNode
	leafs []int
}

func (tree Tree) GetRoot() BNode {
	return tree.root
}

func (tree Tree) GetLeafs() []int {
	return tree.leafs
}
