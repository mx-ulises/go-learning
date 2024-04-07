func AddMask(intersection *map[int]int, nums *[]int, mask int) {
	for _, num := range *nums {
		(*intersection)[num] |= mask
	}
}

func CountElementsInIntersection(intersection *map[int]int, nums *[]int, mask int) int {
	count := 0
	for _, num := range *nums {
		if (*intersection)[num] == mask {
			count++
		}
	}
	return count
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	intersection := map[int]int{}
	AddMask(&intersection, &nums1, 1)
	AddMask(&intersection, &nums2, 2)
	a := CountElementsInIntersection(&intersection, &nums1, 3)
	b := CountElementsInIntersection(&intersection, &nums2, 3)
	return []int{a, b}

}
