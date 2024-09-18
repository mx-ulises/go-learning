func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func GetTime(timePoint string) int {
	h, _ := strconv.Atoi(timePoint[:2])
	m, _ := strconv.Atoi(timePoint[3:])
	return h*60 + m
}

func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	timePointsInt := make([]int, n)
	for i := 0; i < n; i++ {
		timePointsInt[i] = GetTime(timePoints[i])
	}
	sort.Ints(timePointsInt)
	jump := 24 * 60
	minimal := jump
	for i := 0; i < n; i++ {
		j := i + 1
		if j == n {
			j = 0
		}
		candidate := min(abs(timePointsInt[j]-timePointsInt[i]), abs(timePointsInt[j]-timePointsInt[i]+jump))
		minimal = min(minimal, candidate)
	}
	return minimal
}
