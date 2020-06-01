/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxDepth(root *TreeNode) int {
	start:= 0
	return max(countLevel(root, start), countLevel(root, start))
}

func countLevel(node *TreeNode, level int) int {
	if node == nil {
		return level
	}
	level += 1
	return max(countLevel(node.Left, level), countLevel(node.Right, level))
}
// @lc code=end

