func tupleSameProduct(nums []int) int {
	multiples := map[int]int{}
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			multiple := nums[i] * nums[j]
			count += 8 * multiples[multiple]
			multiples[multiple]++
		}
	}
	return count
}
