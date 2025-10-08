package leetcode

func mostFrequentEven(nums []int) int {
	visitCount := map[int]int{}
	mostFrequent, maximalCount := -1, 0
	for _, num := range nums {
		if num&1 == 0 {
			visitCount[num]++
			if maximalCount < visitCount[num] {
				maximalCount = visitCount[num]
				mostFrequent = num
			}
			if visitCount[num] == maximalCount {
				mostFrequent = min(mostFrequent, num)
			}
		}
	}
	return mostFrequent
}
