func numberOfSteps(num int) int {
	steps := 0
	for 0 < num {
		if num&1 == 1 {
			num -= 1
		} else {
			num >>= 1
		}
		steps++
	}
	return steps
}
