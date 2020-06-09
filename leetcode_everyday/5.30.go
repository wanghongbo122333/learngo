package main

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

//判断一棵树是否为镜像树
func isSysmetricTree(root TreeNode) bool {
	return check(&root, &root)
}
func check(Left *TreeNode, Right *TreeNode) bool {
	if Left == nil && Right == nil {
		return true
	}

	if Left == nil || Right == nil {
		return false
	}
	return Left.val == Right.val && check(Left.Left, Right.Right) && check(Left.Right, Right.Left)
}
