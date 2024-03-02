func sortedSquares(nums []int) []int {
	squares := []int{}
	left, right := 0, len(nums)-1
	leftCandidate := nums[left] * nums[left]
	rightCandidate := nums[right] * nums[right]
	for left < right {
		if leftCandidate < rightCandidate {
			squares = append(squares, rightCandidate)
			right--
			rightCandidate = nums[right] * nums[right]
		} else {
			squares = append(squares, leftCandidate)
			left++
			leftCandidate = nums[left] * nums[left]
		}
	}
	squares = append(squares, leftCandidate)
	left, right = 0, len(nums)-1
	for left < right {
		aux := squares[left]
		squares[left] = squares[right]
		squares[right] = aux
		left++
		right--
	}
	return squares
}
