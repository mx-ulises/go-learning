package leetcode

import "sort"

func getColumnClusters(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}
	return matrix
}

func sortRows(matrix [][]int) [][]int {
	for _, row := range matrix {
		sort.Slice(row, func(i, j int) bool {
			return row[i] > row[j]
		})
	}
	return matrix
}

func largestSubmatrix(matrix [][]int) int {
	matrix = getColumnClusters(matrix)
	matrix = sortRows(matrix)
	n, m, maximal := len(matrix), len(matrix[0]), 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			maximal = max((j+1)*matrix[i][j], maximal)
		}
	}
	return maximal
}
