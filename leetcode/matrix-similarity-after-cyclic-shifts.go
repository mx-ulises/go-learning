package leetcode

func getIndex(i, n, k int) int {
	i -= k
	if i < 0 {
		i += n
	} else if n <= i {
		i -= n
	}
	return i
}

func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	k = k % n
	for _, row := range mat {
		for i := 0; i < n; i++ {
			if row[i] != row[getIndex(i, n, k)] {
				return false
			}
		}
		k = -k
	}
	return true
}
