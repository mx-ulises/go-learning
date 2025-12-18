package leetcode

func minCuttingCost(n int, m int, k int) int64 {
	largestLog := max(n, m)
	return getCostToFit(largestLog, k)
}

func getCostToFit(log int, k int) int64 {
	costToFit := int64(k) * int64(log-k)
	return max(0, costToFit)
}
