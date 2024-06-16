func minPatches(nums []int, n int) int {
	miss := 1
	patches := 0
	i := 0
	for miss <= n {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			patches++
		}
	}
	return patches
}
