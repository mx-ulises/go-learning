func GetRelativeOderingMap(arr2 *[]int) *map[int]int {
	relativeOrderingMap := map[int]int{}
	for _, x := range *arr2 {
		relativeOrderingMap[x] = 0
	}
	return &relativeOrderingMap
}

func CountSortAndRemainings(relativeOrderingMap *map[int]int, arr1 *[]int) *[]int {
	remaining := []int{}
	for _, x := range *arr1 {
		_, exists := (*relativeOrderingMap)[x]
		if exists {
			(*relativeOrderingMap)[x]++
		} else {
			remaining = append(remaining, x)
		}
	}
	return &remaining
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	relativeOrderingMap := GetRelativeOderingMap(&arr2)
	remaining := CountSortAndRemainings(relativeOrderingMap, &arr1)
	output := []int{}
	for _, x := range arr2 {
		for i := 0; i < (*relativeOrderingMap)[x]; i++ {
			output = append(output, x)
		}
	}
	sort.Ints(*remaining)
	output = append(output, (*remaining)...)
	return output
}
