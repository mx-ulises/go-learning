package leetcode

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	output := [][]int{}
	minimal := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minimal {
			output = [][]int{}
			minimal = diff
		}
		if diff == minimal {
			output = append(output, []int{arr[i-1], arr[i]})
		}
	}
	return output
}
