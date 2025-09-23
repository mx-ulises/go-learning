package leetcode

func findFinalValue(nums []int, original int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		if num%original == 0 {
			numSet[num] = true
		}
	}
	for numSet[original] {
		original <<= 1
	}
	return original
}
