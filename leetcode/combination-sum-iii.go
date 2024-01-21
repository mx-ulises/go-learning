func BuildCombination(combination *[10]bool) []int {
	element := []int{}
	for i := 1; i < 10; i++ {
		if (*combination)[i] {
			element = append(element, i)
		}
	}
	return element
}

func GetCombinations(start int, k int, n int, combinations *[][]int, combination *[10]bool) {
	if k == 0 && n == 0 {
		(*combinations) = append(*combinations, BuildCombination(combination))
	} else if 0 < k && 0 < n {
		for i := start; i < 10; i++ {
			(*combination)[i] = true
			GetCombinations(i+1, k-1, n-i, combinations, combination)
			(*combination)[i] = false
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	combinations := make([][]int, 0)
	var combination [10]bool
	GetCombinations(1, k, n, &combinations, &combination)
	return combinations
}
