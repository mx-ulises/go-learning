func maximumWealth(accounts [][]int) int {
	output := 0
	for _, account := range accounts {
		customerWealth := 0
		for _, balance := range account {
			customerWealth += balance
		}
		if output < customerWealth {
			output = customerWealth
		}
	}
	return output
}
