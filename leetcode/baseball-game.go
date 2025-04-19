func calPoints(operations []string) int {
	n, i := len(operations), 1
	scores, total := make([]int, n+1), 0
	for _, operation := range operations {
		switch operation {
		case "+":
			scores[i] = scores[i-1] + scores[i-2]
		case "D":
			scores[i] = 2 * scores[i-1]
		case "C":
			total -= scores[i-1] + scores[i-2]
			scores[i-1] = 0
			i -= 2
		default:
			scores[i], _ = strconv.Atoi(operation)
		}
		total += scores[i]
		i++
	}
	return total
}
