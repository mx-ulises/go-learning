package leetcode

func circularGameLosers(n int, k int) []int {
	visited := make([]bool, n)
	i, steps := 0, k
	for !visited[i] {
		visited[i] = true
		i = (i + steps) % n
		steps += k
	}
	losers := []int{}
	for i, v := range visited {
		if !v {
			losers = append(losers, i+1)
		}
	}
	return losers
}
