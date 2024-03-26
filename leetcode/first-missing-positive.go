func firstMissingPositiveMemory(nums []int) int {
	n := 100000
	if len(nums) < n {
		n = len(nums)
	}
	numChecker := make([]bool, n+1)
	for _, num := range nums {
		if 0 < num && num <= n {
			numChecker[num-1] = true
		}
	}
	for i, present := range numChecker {
		if present == false {
			return i + 1
		}
	}
	return -1
}

func GetNum(num, marker int) int {
	for num < 0 {
		num += marker
	}
	return num % marker
}

func firstMissingPositive(nums []int) int {
	start := 1
	end := len(nums)
	marker := end + 1
	// Remove numbers out of bounds
	for i := 0; i < end; i++ {
		if nums[i] < start || end < nums[i] {
			nums[i] = 0
		}
	}
	// Mark visited
	for i := 0; i < end; i++ {
		j := GetNum(nums[i], marker)
		if j != 0 {
			nums[j-1] -= marker
		}
	}
	// Find the first integer not populated
	for i := 0; i < end; i++ {
		if 0 <= nums[i] {
			return i + 1
		}
	}
	return marker
}
