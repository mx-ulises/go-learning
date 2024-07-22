func sum(nums *[]int) int {
	total := 0
	for _, num := range *nums {
		total += num
	}
	return total
}

func addedInteger(nums1 []int, nums2 []int) int {
	return (sum(&nums2) - sum(&nums1)) / len(nums1)
}
