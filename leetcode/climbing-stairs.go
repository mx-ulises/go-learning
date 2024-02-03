func linearMemory(n int) int {
	ways := make([]int, n+2)
	ways[0] = 1
	for i := 0; i < n; i++ {
		ways[i+1] += ways[i]
		ways[i+2] += ways[i]
	}
	return ways[n]
}

func constantMemory(n int) int {
	curr, next, after := 1, 0, 0
	for i := 0; i < n; i++ {
		next += curr
		after += curr
		curr, next, after = next, after, 0
	}
	return curr
}

func climbStairs(n int) int {
	return constantMemory(n)
}
