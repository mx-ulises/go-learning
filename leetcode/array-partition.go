func arrayPairSum(nums []int) int {
	n := len(nums) / 2
	sort.Ints(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i*2]
	}
	return sum
}
