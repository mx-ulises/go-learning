func Sort(intervals *[][]int) {
	sort.Slice((*intervals)[:], func(i int, j int) bool {
		return (*intervals)[i][0] < (*intervals)[j][0]
	})
}

func merge(intervals [][]int) [][]int {
	Sort(&intervals)
	output := make([][]int, 0)
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end && end < intervals[i][1] {
			end = intervals[i][1]
		} else if end < intervals[i][0] {
			output = append(output, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	output = append(output, []int{start, end})
	return output
}
