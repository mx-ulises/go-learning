package leetcode

func digitSum(num int) int {
	sum := 0
	for 0 < num {
		sum += num % 10
		num /= 10
	}
	return sum
}

func smallestIndex(nums []int) int {
	for i, num := range nums {
		if i == digitSum(num) {
			return i
		}
	}
	return -1
}
