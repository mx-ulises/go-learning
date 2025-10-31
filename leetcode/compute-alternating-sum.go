package leetcode

func alternatingSum(nums []int) int {
	total, positive := 0, true
	for _, num := range nums {
		if positive {
			total += num
		} else {
			total -= num
		}
		positive = !positive
	}
	return total
}
