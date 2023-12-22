func GetCombinations(combinations *[][]int, candidates *[]int, preffix *[]int, target int) {
	if target < 0 {
		return
	} else if target == 0 {
		combination := make([]int, 0)
		for i := 0; i < len((*preffix)); i++ {
			combination = append(combination, (*preffix)[i])
		}
		(*combinations) = append((*combinations), combination)
	} else {
		if len(*candidates) == 0 {
			return
		}
		candidate := (*candidates)[len(*candidates)-1]
		(*candidates) = (*candidates)[:len(*candidates)-1]
		total := 0
		for total <= target {
			GetCombinations(combinations, candidates, preffix, target-total)
			total += candidate
			(*preffix) = append((*preffix), candidate)
		}
		removeCount := total / candidate
		(*preffix) = (*preffix)[:len((*preffix))-removeCount]
		(*candidates) = append((*candidates), candidate)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	combinations := make([][]int, 0)
	preffix := make([]int, 0)
	GetCombinations(&combinations, &candidates, &preffix, target)
	return combinations
}
