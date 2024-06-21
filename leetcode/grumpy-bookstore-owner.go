func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	total, unsatisfied := 0, 0
	for i := 0; i < n; i++ {
		total += customers[i]
		if minutes <= i {
			unsatisfied += customers[i] * grumpy[i]
		}
	}
	minimal := unsatisfied
	for i := minutes; i < n; i++ {
		j := i - minutes
		unsatisfied += customers[j] * grumpy[j]
		unsatisfied -= customers[i] * grumpy[i]
		minimal = min(minimal, unsatisfied)
	}
	return total - minimal
}
