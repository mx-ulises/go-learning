func sortColors(nums []int) {
	colorCount := [3]int{}
	for _, color := range nums {
		colorCount[color]++
	}
	currentColor := 0
	for i := 0; i < len(nums); i++ {
		for 0 == colorCount[currentColor] {
			currentColor++
		}
		nums[i] = currentColor
		colorCount[currentColor]--
	}
}
