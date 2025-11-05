package leetcode

func findMissingElements(nums []int) []int {
	found, minimal, maximal := [101]bool{}, 101, 0
	for _, num := range nums {
		minimal = min(minimal, num)
		maximal = max(maximal, num)
		found[num] = true
	}
	missing := []int{}
	for i := minimal; i < maximal; i++ {
		if !found[i] {
			missing = append(missing, i)
		}
	}
	return missing
}
