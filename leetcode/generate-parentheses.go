func InitializeCombinationsMap() map[int]map[string]bool {
	combinations := make(map[int]map[string]bool)
	combinations[1] = make(map[string]bool)
	combinations[1]["()"] = true
	return combinations
}

func BuildCombinations(combinations *map[int]map[string]bool, i int, j int, k int) {
	for combination_j, _ := range (*combinations)[j] {
		for combination_k, _ := range (*combinations)[k] {
			(*combinations)[i][combination_j+combination_k] = true
			if j == 1 {
				(*combinations)[i]["("+combination_k+")"] = true
			}
		}
	}
}

func generateParenthesis(n int) []string {
	combinations := InitializeCombinationsMap()
	for i := 2; i <= n; i++ {
		combinations[i] = make(map[string]bool)
		for j := 1; j < i; j++ {
			BuildCombinations(&combinations, i, j, i-j)
		}
	}
	output := make([]string, 0)
	for combination, _ := range combinations[n] {
		output = append(output, combination)
	}
	return output
}
