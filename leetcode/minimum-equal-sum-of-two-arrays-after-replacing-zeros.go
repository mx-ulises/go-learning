func getSumAndZeros(nums []int) (int, int) {
	sum, zeros := 0, 0
	for _, num := range nums {
		if num == 0 {
			zeros++
		}
		sum += num
	}
	return sum, zeros
}

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, zeros1 := getSumAndZeros(nums1)
	sum2, zeros2 := getSumAndZeros(nums2)
	candidate := max(sum1+zeros1, sum2+zeros2)
	if (zeros1 == 0 && sum1 < candidate) || (zeros2 == 0 && sum2 < candidate) {
		return -1
	}
	return int64(candidate)
}
