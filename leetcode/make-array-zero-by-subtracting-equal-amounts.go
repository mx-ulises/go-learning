package leetcode

func minimumOperations(nums []int) int {
	numSet := map[int]bool{0: true}
	for _, num := range nums {
		numSet[num] = true
	}
	return len(numSet) - 1
}
