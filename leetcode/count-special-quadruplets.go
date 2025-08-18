package leetcode

func countQuadruplets(nums []int) int {
	n, quadruplets := len(nums), 0
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				total := nums[a] + nums[b] + nums[c]
				for d := c + 1; d < n; d++ {
					if nums[d] == total {
						quadruplets++
					}
				}
			}
		}
	}
	return quadruplets
}
