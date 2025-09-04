package leetcode

func isGood(nums []int) bool {
	// We will order the array  by putting the elements in their position
	// To make this simpler, we set the values as 0-base
	for i := 0; i < len(nums); i++ {
		nums[i]--
	}
	// Now the base[n] array will be base[n - 1]
	n := len(nums) - 2
	for i := 0; i < len(nums)-1; i++ {
		// if the number is not in their place, swap it with the one on its position and check again
		for nums[i] != i {
			// Fail if there are numbers out of bounds
			if n < nums[i] {
				return false
			}
			// get position of the value at i
			j := nums[i]
			// if value is n and n is already occupied by n, then just swap with the last element
			if nums[i] == n && nums[n] == n {
				j = n + 1
			}
			// do the swap
			nums[i], nums[j] = nums[j], nums[i]
			// check if after swapping the numbers are equal, if they are, means that there are undesired duplicates and the array is not valid
			if nums[i] == nums[j] {
				return false
			}
		}
	}
	// Check that last element is actually n
	return nums[n+1] == n
}
