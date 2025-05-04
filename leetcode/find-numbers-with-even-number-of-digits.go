func inRange(num, start, end int) bool {
	return start <= num && num <= end
}

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		if inRange(num, 10, 99) || inRange(num, 1000, 9999) || num == 100000 {
			count++
		}
	}
	return count
}
