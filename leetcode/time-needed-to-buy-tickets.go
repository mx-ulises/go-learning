func timeRequiredToBuy(tickets []int, k int) int {
	time, iterations, customers := 0, tickets[k], len(tickets)
	costumersBuyingTickets := make([]int, iterations)
	for i, ticket := range tickets {
		if iterations <= ticket {
			if i <= k {
				time++
			}
		} else {
			costumersBuyingTickets[ticket]++
		}
	}
	for i := 1; i < iterations; i++ {
		time += customers
		customers -= costumersBuyingTickets[i]
	}
	return time
}
