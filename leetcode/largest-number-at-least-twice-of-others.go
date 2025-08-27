package leetcode

func dominantIndex(nums []int) int {
	first, second, output := 0, 0, -1
	for i, num := range nums {
		if first <= num {
			second, first, output = first, num, i
		} else if second < num {
			second = num
		}
	}
	if first < second*2 {
		return -1
	}
	return output
}
