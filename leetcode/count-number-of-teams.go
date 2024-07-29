func ComputeIteration(rating []int, dir int) int {
	total, n := 0, len(rating)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (rating[j] * dir) <= (rating[i] * dir) {
				continue
			}
			for k := j + 1; k < n; k++ {
				if (rating[j] * dir) < (rating[k] * dir) {
					total++
				}
			}
		}
	}
	return total
}

func numTeams(rating []int) int {
	return ComputeIteration(rating, 1) + ComputeIteration(rating, -1)
}
