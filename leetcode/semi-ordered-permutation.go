package leetcode

func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	onePosition := getPosition(nums, 1)
	nPosition := getPosition(nums, n)
	return getTotalSwaps(onePosition, nPosition, n)
}

func getPosition(nums []int, x int) int {
	for i, num := range nums {
		if x == num {
			return i
		}
	}
	return -1
}

func getTotalSwaps(onePosition int, nPosition int, n int) int {
	swaps := getSwaps(onePosition, 0) + getSwaps(nPosition, n-1)
	swaps = adjustSwaps(onePosition, nPosition, swaps)
	return swaps
}

func getSwaps(currentPosition int, targetPosition int) int {
	return abs(targetPosition - currentPosition)
}

func abs(x int) int {
	return max(x, -x)
}

func adjustSwaps(onePosition int, nPosition int, swaps int) int {
	if nPosition < onePosition {
		swaps--
	}
	return swaps
}
