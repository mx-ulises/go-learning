func bisectLeft(nums []int, x int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func bisectRight(nums []int, x int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func maximumCount(nums []int) int {
	negatives := bisectLeft(nums, 0)
	positives := len(nums) - bisectRight(nums, 0)
	return max(negatives, positives)
}
