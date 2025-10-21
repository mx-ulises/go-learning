package leetcode

func minNumber(nums1 []int, nums2 []int) int {
	minimal1, minimal2, minimal := 10, 10, 100
	digitCount := [10]int{}
	for _, num := range nums1 {
		minimal1 = min(minimal1, num)
		digitCount[num]++
	}
	for _, num := range nums2 {
		minimal2 = min(minimal2, num)
		digitCount[num]++
		if digitCount[num] == 2 {
			minimal = min(minimal, num)
		}
	}
	candidate := 10*min(minimal1, minimal2) + max(minimal1, minimal2)
	minimal = min(minimal, candidate)
	return minimal
}
