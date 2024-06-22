func numberOfSubarrays(nums []int, k int) int {
	start, end, n, oddCount := 0, 0, len(nums), 0
	subarrayCount := 0
	for end < n {
		for end < n && oddCount < k {
			if nums[end]&1 == 1 {
				oddCount += 1
			}
			end += 1
		}
		preffixSize := 1
		for end < n && nums[end]&1 == 0 {
			preffixSize += 1
			end += 1
		}
		for start < end && k == oddCount {
			subarrayCount += preffixSize
			if nums[start]&1 == 1 {
				oddCount -= 1
			}
			start += 1
		}
	}
	return subarrayCount
}
