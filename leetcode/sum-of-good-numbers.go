package leetcode

func getVal(nums []int, i int) int {
	if i < 0 || len(nums) <= i {
		return 0
	}
	return nums[i]
}

func sumOfGoodNumbers(nums []int, k int) int {
	sum := 0
	for i, num := range nums {
		if getVal(nums, i-k) < num && getVal(nums, i+k) < num {
			sum += num
		}
	}
	return sum
}
