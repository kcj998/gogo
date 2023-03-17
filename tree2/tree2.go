package main

import "fmt"

// 表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 实现二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	return append(append(left, root.Val), right...)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// 遍历二叉树
	fmt.Println("Inorder Traversal:", inorderTraversal(root))
}
