func pivotArray(nums []int, pivot int) []int {
	left := []int{}
	mid := []int{}
	right := []int{}
	for _, a := range nums {
		if a < pivot {
			left = append(left, a)
		} else if a == pivot {
			mid = append(mid, a)
		} else {
			right = append(right, a)
		}
	}
	left = append(left, mid...)
	left = append(left, right...)
	return left
}
