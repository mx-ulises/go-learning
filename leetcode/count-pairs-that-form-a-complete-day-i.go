package leetcode

func countCompleteDayPairs(hours []int) int {
	visited := map[int]int{}
	count := 0
	for _, h := range hours {
		h = h % 24
		d := (24 - h) % 24
		count += visited[d]
		visited[h]++
	}
	return count
}
