func findScore(nums []int) int64 {
	sortedNums := [][]int{}
	for i, num := range nums {
		sortedNums = append(sortedNums, []int{num, i})
	}
	sort.Slice(sortedNums, func(i, j int) bool {
		if sortedNums[i][0] == sortedNums[j][0] {
			return sortedNums[i][1] < sortedNums[j][1]
		}
		return sortedNums[i][0] < sortedNums[j][0]
	})
	score := 0
	visited := map[int]bool{}
	for _, pair := range sortedNums {
		num, i := pair[0], pair[1]
		if visited[i] == false {
			score += num
			visited[i] = true
			visited[i+1] = true
			visited[i-1] = true
		}
	}
	return int64(score)
}
