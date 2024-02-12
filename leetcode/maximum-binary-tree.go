/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Maximal(nums *[]int, left int, right int) int {
	pivot := left
	for i := left; i <= right; i++ {
		if (*nums)[pivot] < (*nums)[i] {
			pivot = i
		}
	}
	return pivot
}

func ConstructMBT(nums *[]int, left int, right int) *TreeNode {
	if right < left {
		return nil
	} else {
		pivot := Maximal(nums, left, right)
		output := &TreeNode{Val: (*nums)[pivot]}
		output.Left = ConstructMBT(nums, left, pivot-1)
		output.Right = ConstructMBT(nums, pivot+1, right)
		return output
	}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return ConstructMBT(&nums, 0, len(nums)-1)
}
