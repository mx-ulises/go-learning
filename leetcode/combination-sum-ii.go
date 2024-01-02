func GetNumCount(candidates *[]int, target int) *[][]int {
	numCountMap := make(map[int]int)
	for i := 0; i < len(*candidates); i++ {
		num := (*candidates)[i]
		numCountMap[num] += 1
		if target < (numCountMap[num] * num) {
			numCountMap[num] -= 1
		}
	}
	numCount := make([][]int, 0)
	for num, count := range numCountMap {
		if 0 < count {
			numCount = append(numCount, []int{num, count})
		}
	}
	return &numCount
}

func GetCombinations(combinations *[][]int, numCount *[][]int, prefix *[]int, target int) {
	if target == 0 {
		combination := make([]int, len(*prefix))
		copy(combination, *prefix)
		(*combinations) = append((*combinations), combination)
	} else if 0 < target && 0 < len(*numCount) {
		pair := (*numCount)[len(*numCount)-1]
		num, count := pair[0], pair[1]
		(*numCount) = (*numCount)[:len(*numCount)-1]
		for i := 0; i <= count; i++ {
			GetCombinations(combinations, numCount, prefix, target)
			target -= num
			(*prefix) = append(*prefix, num)
		}
		(*prefix) = (*prefix)[:len(*prefix)-count-1]
		(*numCount) = append(*numCount, pair)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	numCount := *(GetNumCount(&candidates, target))
	combinations := make([][]int, 0)
	prefix := []int{}
	GetCombinations(&combinations, &numCount, &prefix, target)
	return combinations
}
