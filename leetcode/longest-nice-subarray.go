func longestNiceSubarray(nums []int) int {
	c, left, maximal := 0, 0, 0
	for right, num := range nums {
		for left < right && (c&num) != 0 {
			c ^= nums[left]
			left++
		}
		maximal = max(maximal, right-left+1)
		c ^= num
	}
	return maximal
}
