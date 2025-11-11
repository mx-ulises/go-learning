package leetcode

import "sort"

func maxDivScore(nums []int, divisors []int) int {
	sort.Ints(nums)
	sort.Ints(divisors)
	maxDivisor, maxScore, start := divisors[0], 0, 0
	for _, divisor := range divisors {
		start = updateStart(nums, divisor, start)
		score := getScore(nums, divisor, start)
		if maxScore < score {
			maxScore = score
			maxDivisor = divisor
		}
	}
	return maxDivisor
}

func updateStart(nums []int, divisor int, start int) int {
	for start < len(nums) && nums[start] < divisor {
		start++
	}
	return start
}

func getScore(nums []int, divisor int, start int) int {
	score := 0
	for i := start; i < len(nums); i++ {
		if nums[i]%divisor == 0 {
			score++
		}
	}
	return score
}
