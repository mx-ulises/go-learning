package leetcode

func calculateTax(brackets [][]int, income int) float64 {
	taxes, prev := 0.0, 0
	for _, bracket := range brackets {
		upper, percent := bracket[0]-prev, float64(bracket[1])/100.0
		taxes += float64(min(income, upper)) * percent
		income, prev = max(0, income-upper), bracket[0]
	}
	return taxes
}
