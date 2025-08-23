package leetcode

func mostFrequent(nums []int, key int) int {
	maxTarget, maxTargetCount, n := 0, 0, len(nums)-1
	targetCount := map[int]int{}
	for i := 0; i < n; i++ {
		if nums[i] == key {
			target := nums[i+1]
			targetCount[target]++
			if maxTargetCount < targetCount[target] {
				maxTargetCount, maxTarget = targetCount[target], target
			}
		}
	}
	return maxTarget
}
