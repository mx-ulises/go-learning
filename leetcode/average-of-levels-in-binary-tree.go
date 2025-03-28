/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getLevelCountAndSum(levelCount, levelSum []int, node *TreeNode, level int) ([]int, []int) {
	if node == nil {
		return levelSum, levelCount
	}
	if len(levelCount) == level {
		levelCount = append(levelCount, 0)
		levelSum = append(levelSum, 0)
	}
	levelSum[level] += node.Val
	levelCount[level]++
	level++
	levelSum, levelCount = getLevelCountAndSum(levelCount, levelSum, node.Left, level)
	levelSum, levelCount = getLevelCountAndSum(levelCount, levelSum, node.Right, level)
	return levelSum, levelCount
}

func averageOfLevels(root *TreeNode) []float64 {
	levelCount, levelSum := []int{}, []int{}
	levelSum, levelCount = getLevelCountAndSum(levelCount, levelSum, root, 0)
	output := make([]float64, len(levelCount))
	for i := 0; i < len(output); i++ {
		output[i] = float64(levelSum[i]) / float64(levelCount[i])
	}
	return output
}
