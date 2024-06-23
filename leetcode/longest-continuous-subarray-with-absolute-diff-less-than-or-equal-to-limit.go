func longestSubarray(nums []int, limit int) int {
	start, end, n := 0, 0, len(nums)
	minimal, maximal := []int{}, []int{}
	longest := 0
	for end < n {
		for 0 < len(minimal) && nums[end] < minimal[len(minimal)-1] {
			minimal = minimal[0 : len(minimal)-1]
		}
		minimal = append(minimal, nums[end])
		for 0 < len(maximal) && maximal[len(maximal)-1] < nums[end] {
			maximal = maximal[0 : len(maximal)-1]
		}
		maximal = append(maximal, nums[end])
		for start < end && limit < (maximal[0]-minimal[0]) {
			if minimal[0] == nums[start] {
				minimal = minimal[1:]
			}
			if maximal[0] == nums[start] {
				maximal = maximal[1:]
			}
			start += 1
		}
		end += 1
		longest = max(longest, end-start)
	}
	return longest
}
