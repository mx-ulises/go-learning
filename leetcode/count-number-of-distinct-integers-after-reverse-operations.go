func reverse(num int) int {
	newNum := 0
	for 0 < num {
		newNum = (newNum * 10) + (num % 10)
		num /= 10
	}
	return newNum
}

func countDistinctIntegers(nums []int) int {
	distinctInts := map[int]bool{}
	for _, num := range nums {
		distinctInts[num] = true
		distinctInts[reverse(num)] = true
	}
	return len(distinctInts)
}
