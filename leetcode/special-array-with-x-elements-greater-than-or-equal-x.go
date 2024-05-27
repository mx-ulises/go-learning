func specialArray(nums []int) int {
	sort.Ints(nums)
	i, n, prev := 0, len(nums), -1
	for i < n {
		// Get value of smallest remaining items
		aux := nums[i]
		// Get count of remaining items
		x := n - i
		// Previous item value is larger than count of remaining values of the list, so we cannot continue
		if x <= prev {
			return -1
		}
		// Value of remaining items is larger than the count of remaining items, we can return this value
		if x <= aux {
			return x
		}
		// Remove all items that are equals to the current reviewed number
		for i < n && nums[i] == aux {
			i++
		}
		// Update previous value
		prev = aux
	}
	// List traversed without success
	return -1
}
