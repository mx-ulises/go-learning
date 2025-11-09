package leetcode

func findClosestNumber(nums []int) int {
	closest, closestDistance := 1000000, 1000000
	for _, num := range nums {
		distance := getDistance(num)
		if distance < closestDistance {
			closestDistance = distance
			closest = num
		} else if distance == closestDistance {
			closest = max(closest, num)
		}
	}
	return closest
}

func getDistance(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
