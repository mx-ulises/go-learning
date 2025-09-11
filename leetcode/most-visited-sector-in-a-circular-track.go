package leetcode

import "sort"

func mostVisited(n int, rounds []int) []int {
	start := rounds[0]
	end := rounds[len(rounds)-1] - start
	if end < 0 {
		end += n
	}
	output := []int{}
	for i := 0; i <= end; i++ {
		num := (i + start) % n
		if num == 0 {
			num = n
		}
		output = append(output, num)
	}
	sort.Ints(output)
	return output
}
