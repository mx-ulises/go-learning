func sort_diagonal(mat [][]int, x, y, n, m int) {
	a := []int{}
	i, j := x, y
	for i < n && j < m {
		a = append(a, mat[i][j])
		i, j = i+1, j+1
	}
	sort.Ints(a)
	i, j, k := x, y, 0
	for i < n && j < m {
		mat[i][j] = a[k]
		i, j, k = i+1, j+1, k+1
	}
}

func diagonalSort(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	sort_diagonal(mat, 0, 0, n, m)
	for i := 1; i < n; i++ {
		sort_diagonal(mat, i, 0, n, m)
	}
	for j := 1; j < m; j++ {
		sort_diagonal(mat, 0, j, n, m)
	}
	return mat
}
