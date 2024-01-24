/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetOddDifference(digitCount *[10]int, digit int) int {
	if (*digitCount)[digit]&1 == 1 {
		return 1
	}
	return -1
}

func GetPseudoPalindromicPaths(node *TreeNode, digitCount *[10]int, odds int) int {
	if node == nil {
		return 0
	}
	(*digitCount)[node.Val]++
	oddDirection := GetOddDifference(digitCount, node.Val)
	odds += oddDirection
	totalPaths := 0
	if node.Left == nil && node.Right == nil && odds < 2 {
		totalPaths = 1
	}
	if node.Left != nil {
		totalPaths += GetPseudoPalindromicPaths(node.Left, digitCount, odds)
	}
	if node.Right != nil {
		totalPaths += GetPseudoPalindromicPaths(node.Right, digitCount, odds)
	}
	(*digitCount)[node.Val]--
	return totalPaths
}

func pseudoPalindromicPaths(root *TreeNode) int {
	digitCount := &[10]int{}
	return GetPseudoPalindromicPaths(root, digitCount, 0)
}
