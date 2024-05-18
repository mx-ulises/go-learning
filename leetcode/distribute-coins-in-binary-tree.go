/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Traverse(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	coinsLeft, movesLeft := Traverse(node.Left)
	coinsRight, movesRight := Traverse(node.Right)
	coins := coinsLeft + coinsRight + node.Val - 1
	moves := movesLeft + movesRight + abs(coinsLeft) + abs(coinsRight)
	return coins, moves
}

func distributeCoins(root *TreeNode) int {
	_, moves := Traverse(root)
	return moves
}
