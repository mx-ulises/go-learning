package leetcode

func missingMultiple(nums []int, k int) int {
	visited := map[int]bool{}
	for _, num := range nums {
		if num%k == 0 {
			visited[num] = true
		}
	}
	missing := k
	for visited[missing] {
		missing += k
	}
	return missing
}
