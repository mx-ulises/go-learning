package leetcode

import "sort"

func absDifference(nums []int, k int) int {
	sort.Ints(nums)
	largestKSum := getLargestKSum(nums, k)
	shortestKSum := getShortestKSum(nums, k)
	return largestKSum - shortestKSum
}

func getLargestKSum(nums []int, k int) int {
	start := len(nums) - k
	end := len(nums)
	return getSumFromTo(nums, start, end)
}

func getShortestKSum(nums []int, k int) int {
	start := 0
	end := k
	return getSumFromTo(nums, start, end)
}

func getSumFromTo(nums []int, start int, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		sum += nums[i]
	}
	return sum
}
