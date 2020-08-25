package binsearchtree

type opFunc func(*binaryTreeNode)

type binaryTreeNode struct {
	val   int
	left  *binaryTreeNode
	right *binaryTreeNode
}

type binaryTree struct {
	root  *binaryTreeNode
	_size int
}

func newBinaryTree() *binaryTree {
	var tree binaryTree
	tree.root = nil
	tree._size = 0
	return &tree
}

func _preorderTraverse(node *binaryTreeNode, op opFunc) {
	if node == nil {
		return
	}
	op(node)
	_preorderTraverse(node.left, op)
	_preorderTraverse(node.right, op)
}

func _midorderTraverse(node *binaryTreeNode, op opFunc) {
	if node == nil {
		return
	}
	_midorderTraverse(node.left, op)
	op(node)
	_midorderTraverse(node.right, op)
}

func _postorderTraverse(node *binaryTreeNode, op opFunc) {
	if node == nil {
		return
	}
	_postorderTraverse(node.left, op)
	_postorderTraverse(node.right, op)
	op(node)
}

func (tree *binaryTree) preorderTraverse(op opFunc) {
	_preorderTraverse(tree.root, op)
}

func (tree *binaryTree) midorderTraverse(op opFunc) {
	_midorderTraverse(tree.root, op)
}

func (tree *binaryTree) postorderTraverse(op opFunc) {
	_postorderTraverse(tree.root, op)
}

func (tree *binaryTree) size() int {
	return tree._size
}

func (tree *binaryTree) isEmpty() bool {
	return tree.size() == 0
}
