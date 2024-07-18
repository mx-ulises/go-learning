/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetGoodLeafPairs(leftLeaves, rightLeaves *map[int]int, distance, level int) int {
	count := 0
	for leftLevel, leftCount := range *leftLeaves {
		for rightLevel, rightCount := range *rightLeaves {
			leavesDistance := leftLevel + rightLevel - 2*level
			if leavesDistance <= distance {
				count += leftCount * rightCount
			}
		}
	}
	return count
}

func CombineLeaves(leftLeaves, rightLeaves *map[int]int) *map[int]int {
	leaves := rightLeaves
	for level, count := range *leftLeaves {
		(*leaves)[level] += count
	}
	return leaves
}

func TraverseCountPairs(node *TreeNode, distance, level int) (*map[int]int, int) {
	if node == nil {
		return &map[int]int{}, 0
	}
	if node.Left == nil && node.Right == nil {
		return &map[int]int{level: 1}, 0
	}
	leftLeaves, leftCount := TraverseCountPairs(node.Left, distance, level+1)
	rightLeaves, rightCount := TraverseCountPairs(node.Right, distance, level+1)
	count := GetGoodLeafPairs(leftLeaves, rightLeaves, distance, level) + leftCount + rightCount
	leaves := CombineLeaves(leftLeaves, rightLeaves)
	return leaves, count
}

func countPairs(root *TreeNode, distance int) int {
	_, count := TraverseCountPairs(root, distance, 0)
	return count
}
