func check(nums []int) bool {
	n, errors := len(nums), 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		if nums[j] < nums[i] {
			errors++
		}
	}
	return errors <= 1
}
