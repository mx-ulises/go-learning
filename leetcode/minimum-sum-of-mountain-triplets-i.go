package leetcode

func minimumSum(nums []int) int {
	minimal := 150
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] <= nums[i] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[j] <= nums[k] {
					continue
				}
				minimal = min(minimal, nums[i]+nums[j]+nums[k])
			}
		}
	}
	if minimal == 150 {
		minimal = -1
	}
	return minimal
}
