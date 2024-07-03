func GetMinDiff(nums *[]int, start int, end int, moves int) int {
	if moves == 0 || start == end {
		return (*nums)[end] - (*nums)[start]
	}
	return min(GetMinDiff(nums, start+1, end, moves-1), GetMinDiff(nums, start, end-1, moves-1))
}

func minDifference(nums []int) int {
	start, end := 0, len(nums)-1
	sort.Ints(nums)
	return GetMinDiff(&nums, start, end, 3)
}
