type LandDescription struct {
	IsFarmLand bool
	X          int
	Y          int
}

func IsFarmLand(land *[][]int, i, j, n, m int) LandDescription {
	if i == n || j == m || (*land)[i][j] == 0 {
		return LandDescription{false, -1, -1}
	}
	(*land)[i][j] = 0
	right := IsFarmLand(land, i+1, j, n, m)
	down := IsFarmLand(land, i, j+1, n, m)
	i = max(i, right.X, down.X)
	j = max(j, right.Y, down.Y)
	return LandDescription{true, i, j}
}

func findFarmland(land [][]int) [][]int {
	farmlands := [][]int{}
	n, m := len(land), len(land[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			desc := IsFarmLand(&land, i, j, n, m)
			if desc.IsFarmLand {
				farmlands = append(farmlands, []int{i, j, desc.X, desc.Y})
			}
		}
	}
	return farmlands
}
