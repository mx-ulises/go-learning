func checkSubarraySum(nums []int, k int) bool {
	sum := 0
	memory := map[int]int{0: 0}
	for current := 0; current < len(nums); current++ {
		sum += nums[current]
		sum %= k
		prev, exists := memory[sum]
		if exists && 1 <= current-prev {
			return true
		}
		if exists == false {
			memory[sum] = current + 1
		}
	}
	return false
}
