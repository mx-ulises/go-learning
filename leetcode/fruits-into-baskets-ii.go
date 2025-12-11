package leetcode

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	unnplacedFruits := 0
	for _, fruit := range fruits {
		if !checkAndPlaceFruit(fruit, baskets) {
			unnplacedFruits++
		}
	}
	return unnplacedFruits
}

func checkAndPlaceFruit(fruit int, baskets []int) bool {
	for i, basket := range baskets {
		if fruit <= basket {
			baskets[i] = 0
			return true
		}
	}
	return false
}
