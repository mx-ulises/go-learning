package leetcode

func minimumRightShifts(nums []int) int {
	start := getStartIndex(nums)
	if !isValidArray(nums, start) {
		return -1
	}
	return len(nums) - start
}

func getStartIndex(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return i
		}
	}
	return len(nums)
}

func isValidArray(nums []int, start int) bool {
	n := len(nums)
	breaches := 0
	for i := 0; i < n; i++ {
		current, succesor := computeCurrentAndSuccesor(i, start, n)
		breaches = updateBreaches(breaches, current, succesor, nums)
	}
	return breaches < 2
}

func computeCurrentAndSuccesor(i int, start int, n int) (int, int) {
	current := (i + start) % n
	succesor := (i + start + 1) % n
	return current, succesor
}

func updateBreaches(breaches int, current int, succesor int, nums []int) int {
	if nums[succesor] < nums[current] {
		breaches++
	}
	return breaches
}
