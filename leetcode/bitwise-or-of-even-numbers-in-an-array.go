package leetcode

func evenNumberBitwiseORs(nums []int) int {
	result := 0
	for _, num := range nums {
		if num&1 == 0 {
			result |= num
		}
	}
	return result
}
