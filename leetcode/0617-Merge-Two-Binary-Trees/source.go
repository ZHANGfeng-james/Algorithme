package leetcode

import (
	"fmt"
)

const (
	topic = "Leetcode Problem 0461. 汉明距离\n"
)

func init() {
	fmt.Println(topic)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 != nil && root2 == nil {
		return root1
	}

	if root1 == nil && root2 != nil {
		return root2
	}

	root1.Val += root2.Val

	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}
