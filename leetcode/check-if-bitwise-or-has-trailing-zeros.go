package leetcode

func hasTrailingZeros(nums []int) bool {
	evenFound := false
	for _, num := range nums {
		if num&1 == 0 {
			if evenFound {
				return true
			}
			evenFound = true
		}
	}
	return false
}
