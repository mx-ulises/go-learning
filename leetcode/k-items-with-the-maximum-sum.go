package leetcode

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	return min(k, numOnes) - max(0, k-numOnes-numZeros)
}
