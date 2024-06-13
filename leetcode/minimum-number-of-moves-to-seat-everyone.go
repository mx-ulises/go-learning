func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	n := len(seats)
	moves := 0
	for i := 0; i < n; i++ {
		moves += abs(seats[i] - students[i])
	}
	return moves
}
