func trap(height []int) int {
	left, right := 0, len(height)-1
	totalArea, filledArea := 0, 0
	prevMin := 0
	for left <= right {
		currentMin := min(height[left], height[right])
		totalArea += (right - left + 1) * (currentMin - prevMin)
		prevMin = currentMin
		for left <= right && height[left] <= currentMin {
			filledArea += height[left]
			left++
		}
		for left <= right && height[right] <= currentMin {
			filledArea += height[right]
			right--
		}
	}
	return totalArea - filledArea
}
