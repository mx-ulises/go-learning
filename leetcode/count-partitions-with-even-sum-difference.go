func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func countPartitions(nums []int) int {
	left, right := 0, sum(nums)
	n := len(nums) - 1
	partitions := 0
	for i := 0; i < n; i++ {
		left += nums[i]
		right -= nums[i]
		diff := left - right
		if (diff & 1) == 0 {
			partitions++
		}
	}
	return partitions
}
