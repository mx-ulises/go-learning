func flipAndInvertImage(image [][]int) [][]int {
	n, m := len(image), len(image[0])
	for i := 0; i < n; i++ {
		// Reverse
		for j := 0; j < m/2; j++ {
			k := m - j - 1
			aux := image[i][j]
			image[i][j] = image[i][k]
			image[i][k] = aux
		}
		// Invert
		for j := 0; j < m; j++ {
			image[i][j] ^= 1
		}
	}
	return image
}
