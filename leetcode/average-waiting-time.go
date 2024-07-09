func averageWaitingTime(customers [][]int) float64 {
	currentTime := 0
	totalWaitingTime := 0
	for _, customer := range customers {
		currentTime = max(currentTime, customer[0]) + customer[1]
		totalWaitingTime += currentTime - customer[0]
	}
	return float64(totalWaitingTime) / float64(len(customers))
}
