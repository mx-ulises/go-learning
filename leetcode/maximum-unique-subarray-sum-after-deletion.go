package leetcode

func maxSum(nums []int) int {
	visited, total, maximal := [101]bool{}, 0, -101
	for _, num := range nums {
		if 0 <= num && !visited[num] {
			total += num
			visited[num] = true
		}
		maximal = max(maximal, num)
	}
	if total == 0 && !visited[0] {
		total = maximal
	}
	return total
}
