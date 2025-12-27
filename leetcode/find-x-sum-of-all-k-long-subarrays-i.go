package leetcode

import "sort"

func findXSum(nums []int, k int, x int) []int {
	buffer := initializeBuffer(nums, k)
	xSum := []int{getXSum(buffer, x)}
	for i := 0; i < len(nums)-k; i++ {
		buffer = updateBuffer(buffer, nums[i], nums[i+k])
		xSum = append(xSum, getXSum(buffer, x))
	}
	return xSum
}

func initializeBuffer(nums []int, k int) map[int]int {
	buffer := map[int]int{}
	for i := 0; i < k; i++ {
		buffer[nums[i]]++
	}
	return buffer
}

func getXSum(buffer map[int]int, x int) int {
	pairs := getPairs(buffer)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] > pairs[j][1]
		}
		return pairs[i][0] > pairs[j][0]
	})
	return sumTopX(pairs, x)
}

func getPairs(buffer map[int]int) [][2]int {
	pairs := [][2]int{}
	for num, count := range buffer {
		pairs = append(pairs, [2]int{count, num})
	}
	return pairs
}

func sumTopX(pairs [][2]int, x int) int {
	sum := 0
	for i := 0; i < min(len(pairs), x); i++ {
		sum += pairs[i][0] * pairs[i][1]
	}
	return sum
}

func updateBuffer(buffer map[int]int, remove int, add int) map[int]int {
	buffer[remove]--
	buffer[add]++
	return buffer
}
