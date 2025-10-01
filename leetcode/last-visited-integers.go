package leetcode

func lastVisitedIntegers(nums []int) []int {
	seen, ans, k := []int{}, []int{}, 0
	for _, num := range nums {
		if 0 < num {
			seen = append(seen, num)
			k = 0
		} else {
			k++
			if k <= len(seen) {
				ans = append(ans, seen[len(seen)-k])
			} else {
				ans = append(ans, -1)
			}
		}
	}
	return ans
}
