func GetSortedNums(nums *[]int, n int) *[]int {
	sorted_nums := make([]int, n)
	for i := 0; i < n; i++ {
		sorted_nums[i] = (*nums)[i]
	}
	sort.Ints(sorted_nums)
	return &sorted_nums
}

func GetPosition(sorted_nums *[]int, n, num int) int {
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if (*sorted_nums)[mid] < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	sorted_nums := GetSortedNums(&nums, n)
	output := make([]int, n)
	for i, num := range nums {
		output[i] = GetPosition(sorted_nums, n, num)
	}
	return output
}
