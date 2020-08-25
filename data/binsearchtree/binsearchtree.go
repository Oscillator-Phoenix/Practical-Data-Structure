package binsearchtree

type lessFunc func(x, y int) bool

type binarySearchTree struct {
	binaryTree
	less lessFunc
}

func newBinarySearchTree(less lessFunc) *binarySearchTree {
	var tree binarySearchTree
	tree.root = nil
	tree._size = 0
	tree.less = less
	return &tree
}

func _insert(node *binaryTreeNode, val int, less lessFunc) *binaryTreeNode {
	if node == nil {
		return &binaryTreeNode{
			val:   val,
			left:  nil,
			right: nil,
		}
	}

	if less(val, node.val) {
		node.left = _insert(node.left, val, less) // insert to left subtree
		return node
	}

	node.right = _insert(node.right, val, less) // insert to right subtree
	return node
}

func (tree *binarySearchTree) insert(val int) {
	tree.root = _insert(tree.root, val, tree.less)
}

func _find(node *binaryTreeNode, val int, less lessFunc) (ok bool) {
	if node == nil {
		return false
	}

	if node.val == val {
		return true
	}

	if less(val, node.val) {
		return _find(node.left, val, less) // insert to left subtree
	}

	return _find(node.right, val, less) // insert to right subtree
}

func (tree *binarySearchTree) find(val int) (ok bool) {
	return _find(tree.root, val, tree.less)
}

func (tree *binarySearchTree) sortedVals() []int {
	sorted := []int{}
	tree.midorderTraverse(func(node *binaryTreeNode) {
		sorted = append(sorted, node.val)
	})
	return sorted
}
