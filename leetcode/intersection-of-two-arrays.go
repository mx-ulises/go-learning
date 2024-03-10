func intersection(nums1 []int, nums2 []int) []int {
	numSet := map[int]bool{}
	for _, x := range nums1 {
		numSet[x] = true
	}
	output := []int{}
	for _, x := range nums2 {
		if numSet[x] {
			output = append(output, x)
			numSet[x] = false
		}
	}
	return output
}
