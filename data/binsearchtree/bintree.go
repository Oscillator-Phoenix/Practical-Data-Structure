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

func maxTwoInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxInts(xs ...int) int {
	if len(xs) == 1 {
		return xs[0]
	}
	return maxTwoInt(xs[0], maxInts(xs[1:]...))
}

func _depth(node *binaryTreeNode) int {
	if node.left == nil && node.right == nil {
		return 1
	}
	if node.left != nil && node.right == nil {
		return 1 + _depth(node.left)
	}
	if node.left == nil && node.right != nil {
		return 1 + _depth(node.right)
	}
	return 1 + maxInts(_depth(node.left), _depth(node.right))
}

func (tree *binaryTree) depth() int {
	return _depth(tree.root)
}

func (tree *binaryTree) levelTraverseBFS(op opFunc) {
	q := [](*binaryTreeNode){}

	if tree.root != nil {
		q = append(q, tree.root) // push
	}

	for len(q) > 0 {
		nq := [](*binaryTreeNode){}

		for len(q) > 0 {
			if q[0].left != nil {
				nq = append(nq, q[0].left)
			}
			if q[0].right != nil {
				nq = append(nq, q[0].right)
			}

			op(q[0])
			q = q[1:] // pop
		}

		q = nq
	}
}

func (tree *binaryTree) levelTraverseDFS() [][]int {
	res := [][]int{}

	var helper func(node *binaryTreeNode, level int)

	helper = func(node *binaryTreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.val)
		helper(node.left, level+1)
		helper(node.right, level+1)
	}

	helper(tree.root, 0)

	return res
}
