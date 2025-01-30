func sumOfUnique(nums []int) int {
	numCount := map[int]int{}
	sum := 0
	for _, num := range nums {
		numCount[num]++
		switch numCount[num] {
		case 1:
			sum += num
		case 2:
			sum -= num
		}
	}
	return sum
}
