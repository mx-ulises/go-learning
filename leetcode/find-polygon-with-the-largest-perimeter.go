func largestPerimeter(nums []int) int64 {
	slices.Sort(nums)
	perimeter := int64(-1)
	currentSum := int64(nums[0]) + int64(nums[1])
	for i := 2; i < len(nums); i++ {
		num := int64(nums[i])
		if num < currentSum {
			perimeter = num + currentSum
		}
		currentSum += num
	}
	return perimeter
}
