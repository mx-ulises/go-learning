func GetStart(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if mid == 0 && nums[mid] == target {
			return mid
		} else if nums[mid] == target && nums[mid-1] != target {
			return mid
		} else if nums[mid] == target && nums[mid-1] == target {
			end = mid - 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func GetEnd(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if mid == (len(nums)-1) && nums[mid] == target {
			return mid
		} else if nums[mid] == target && nums[mid+1] != target {
			return mid
		} else if nums[mid] == target && nums[mid+1] == target {
			start = mid + 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	output := make([]int, 0)
	output = append(output, GetStart(nums, target))
	output = append(output, GetEnd(nums, target))
	return output
}
