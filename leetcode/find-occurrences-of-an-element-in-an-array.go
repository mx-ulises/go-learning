package leetcode

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	n := 0
	for i, num := range nums {
		if num == x {
			nums[n] = i
			n++
		}
	}
	for i, q := range queries {
		q--
		if n <= q {
			queries[i] = -1
		} else {
			queries[i] = nums[q]
		}
	}
	return queries
}
