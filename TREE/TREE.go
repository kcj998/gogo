package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val) // 先输出根节点的值
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 中序遍历
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Println(root.Val) // 在遍历左子树后输出根节点的值
	inorderTraversal(root.Right)
}

// 后序遍历
func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Println(root.Val) // 最后输出根节点的值
}

func main() {
	// 创建一个二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// 先序遍历
	fmt.Println("先序遍历结果：")
	preorderTraversal(root)

	// 中序遍历
	fmt.Println("中序遍历结果：")
	inorderTraversal(root)

	// 后序遍历
	fmt.Println("后序遍历结果：")
	postorderTraversal(root)
}
