func smallestEqual(nums []int) int {
	output := -1
	for i, num := range nums {
		if i%10 == num {
			return i
		}
	}
	return output
}
