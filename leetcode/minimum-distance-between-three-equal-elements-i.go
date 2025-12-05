package leetcode

func minimumDistance(nums []int) int {
	minimalDistance := 300
	numIndexes := [101][]int{}
	for i, num := range nums {
		numIndexes[num] = append(numIndexes[num], i)
		minimalDistance = checkAndUpdateMinimalDistance(numIndexes[num], minimalDistance)
	}
	return validateMinimalDistance(minimalDistance)
}

func checkAndUpdateMinimalDistance(indexes []int, minimalDistance int) int {
	if 3 <= len(indexes) {
		minimalDistance = min(minimalDistance, getDistanceFromLastThreeIndexes(indexes))
	}
	return minimalDistance
}

func getDistanceFromLastThreeIndexes(indexes []int) int {
	i, j, k := getLastThreeIndexes(indexes)
	return abs(i-j) + abs(j-k) + abs(k-i)
}

func getLastThreeIndexes(indexes []int) (int, int, int) {
	n := len(indexes)
	return indexes[n-3], indexes[n-2], indexes[n-1]
}

func abs(x int) int {
	return max(x, -x)
}

func validateMinimalDistance(minimalDistance int) int {
	if minimalDistance == 300 {
		minimalDistance = -1
	}
	return minimalDistance
}
