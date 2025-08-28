package leetcode

func isPrime(x int, visited map[int]bool) bool {
	_, ok := visited[x]
	if ok {
		return visited[x]
	}
	mid := x / 2
	visited[x] = true
	for i := 2; i <= mid; i++ {
		if x%i == 0 {
			visited[x] = false
			break
		}
	}
	return visited[x]
}

func diagonalPrime(nums [][]int) int {
	visited := map[int]bool{0: false, 1: false}
	maximalPrime, n := 0, len(nums)
	for i := 0; i < n; i++ {
		j := n - i - 1
		if isPrime(nums[i][i], visited) {
			maximalPrime = max(maximalPrime, nums[i][i])
		}
		if isPrime(nums[i][j], visited) {
			maximalPrime = max(maximalPrime, nums[i][j])
		}
	}
	return maximalPrime
}
