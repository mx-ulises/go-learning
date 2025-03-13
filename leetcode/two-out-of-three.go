func addNumMask(numMask map[int]int, nums []int, mask int) {
	for _, num := range nums {
		numMask[num] |= mask
	}
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	numMask := map[int]int{}
	addNumMask(numMask, nums1, 1)
	addNumMask(numMask, nums2, 2)
	addNumMask(numMask, nums3, 4)
	output := []int{}
	invalidMasks := map[int]bool{1: true, 2: true, 4: true}
	for num, mask := range numMask {
		if !invalidMasks[mask] {
			output = append(output, num)
		}
	}
	return output
}
