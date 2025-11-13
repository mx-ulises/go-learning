package leetcode

func minOperations(nums []int, k int) int {
	visited, saw := [51]bool{}, 0
	for i := len(nums) - 1; 0 <= i; i-- {
		num := nums[i]
		if num <= k && !visited[num] {
			visited[num] = true
			saw++
		}
		if saw == k {
			return len(nums) - i
		}
	}
	return -1
}
