package leetcode

func unequalTriplets(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if nums[i] != nums[k] && nums[j] != nums[k] {
					count++
				}
			}
		}
	}
	return count
}
