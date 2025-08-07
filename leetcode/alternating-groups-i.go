package leetcode

func numberOfAlternatingGroups(colors []int) int {
	count, n := 0, len(colors)
	prev := colors[n-1]
	for i := 0; i < (n - 1); i++ {
		if prev != colors[i] && colors[i] != colors[i+1] {
			count++
		}
		prev = colors[i]
	}
	if prev != colors[n-1] && colors[n-1] != colors[0] {
		count++
	}
	return count
}
