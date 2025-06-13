func targetIndices(nums []int, target int) []int {
	targetOcurrence, targetPreffixSize := 0, 0
	for _, num := range nums {
		if num < target {
			targetPreffixSize++
		}
		if num == target {
			targetOcurrence++
		}
	}
	output := make([]int, targetOcurrence)
	for i := 0; i < targetOcurrence; i++ {
		output[i] = targetPreffixSize
		targetPreffixSize++
	}
	return output
}
