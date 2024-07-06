func passThePillow(n int, time int) int {
	current, direction, maxPass := 1, 1, n-1
	for 0 < time {
		pass := min(maxPass, time)
		current += direction * pass
		time -= pass
		direction = -direction
	}
	return current
}
