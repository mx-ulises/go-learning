func AddNums(nums *map[int]int, toBeAdded *[]int, mask int) {
	for _, num := range *toBeAdded {
		(*nums)[num] |= mask
	}
}

func AddDistinct(output *[]int, mask, expected_mask, num int) {
	if mask == expected_mask {
		(*output) = append(*output, num)
	}
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums := map[int]int{}
	AddNums(&nums, &nums1, 1)
	AddNums(&nums, &nums2, 2)
	output := make([][]int, 2)
	for num, mask := range nums {
		AddDistinct(&(output[0]), mask, 1, num)
		AddDistinct(&(output[1]), mask, 2, num)
	}
	return output
}
