func GetCount(nums *[]int) *map[int]int {
	count := map[int]int{}
	for _, num := range *nums {
		count[num]++
	}
	return &count
}

func intersect(nums1 []int, nums2 []int) []int {
	count1 := GetCount(&nums1)
	count2 := GetCount(&nums2)
	intersect := []int{}
	for num := range *count1 {
		count := min((*count1)[num], (*count2)[num])
		for i := 0; i < count; i++ {
			intersect = append(intersect, num)
		}
	}
	return intersect
}
