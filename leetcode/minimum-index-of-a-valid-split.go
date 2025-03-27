func getDominant(nums []int) int {
	dominant, dominantCount := 0, 0
	for _, num := range nums {
		if dominant == 0 {
			dominant = num
			dominantCount = 1
		} else if dominant == num {
			dominantCount++
		} else {
			dominantCount--
			if dominantCount == 0 {
				dominant = 0
			}
		}
	}
	return dominant
}

func getCount(nums []int, dominant int) int {
	dominantCount := 0
	for _, num := range nums {
		if num == dominant {
			dominantCount++
		}
	}
	return dominantCount
}

func isDominant(total, dominantCount int) bool {
	return (total - dominantCount) < dominantCount
}

func minimumIndex(nums []int) int {
	dominant := getDominant(nums)
	rightDominantCount := getCount(nums, dominant)
	leftDominantCount, n := 0, len(nums)
	for i, num := range nums {
		if num == dominant {
			leftDominantCount++
			rightDominantCount--
		}
		totalLeft := i + 1
		totalRight := n - totalLeft
		if isDominant(totalLeft, leftDominantCount) && isDominant(totalRight, rightDominantCount) {
			return i
		}
	}
	return -1
}
