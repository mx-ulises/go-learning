/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func allPossibleFBT(n int) []*TreeNode {
	if n&1 == 0 {
		return nil
	}
	subtrees := map[int][]*TreeNode{1: {&TreeNode{}}}
	i := 3
	for i <= n {
		subtrees[i] = []*TreeNode{}
		leftIndex, rigthIndex := 1, i-2
		for 0 < rigthIndex {
			for _, left := range subtrees[leftIndex] {
				for _, right := range subtrees[rigthIndex] {
					subtrees[i] = append(subtrees[i], &TreeNode{Left: left, Right: right})
				}
			}
			leftIndex += 2
			rigthIndex -= 2
		}
		i += 2
	}
	return subtrees[n]
}
