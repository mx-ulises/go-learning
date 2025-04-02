func maximumTripletValue(nums []int) int64 {
	maximal := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				maximal = max(maximal, (nums[i]-nums[j])*nums[k])
			}
		}
	}
	return int64(maximal)
}
