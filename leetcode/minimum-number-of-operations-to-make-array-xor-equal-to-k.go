func minOperations(nums []int, k int) int {
	for _, num := range nums {
		k ^= num
	}
	count := 0
	for 0 < k {
		count += k & 1
		k = k >> 1
	}
	return count
}
