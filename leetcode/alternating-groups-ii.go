func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	start, end, m := 0, 1, n+k-1
	count := 0
	for end < m {
		for end < m && colors[(end-1)%n] != colors[end%n] {
			end++
		}
		count += max(0, end-start-k+1)
		start = end
		end++
	}
	return count
}
