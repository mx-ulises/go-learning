func Sum(array []int) int {
	total := 0
	for _, x := range array {
		total += x
	}
	return total
}

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	total := mean * (m + n)
	pending := max(total-Sum(rolls), 0)
	minimal := pending / n
	maximal := minimal + 1
	maximal_count := pending % n
	minimal_count := n - maximal_count
	missing := make([]int, n)
	for i := 0; i < minimal_count; i++ {
		missing[i] = minimal
	}
	for i := minimal_count; i < n; i++ {
		missing[i] = maximal
	}
	if missing[0] < 1 || 6 < missing[n-1] {
		return []int{}
	}
	return missing
}
