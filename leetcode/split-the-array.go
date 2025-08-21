package leetcode

func isPossibleToSplit(nums []int) bool {
	numCount := [101]int{}
	for _, num := range nums {
		numCount[num]++
		if numCount[num] == 3 {
			return false
		}
	}
	return true
}
