package leetcode

func sumDivisibleByK(nums []int, k int) int {
	frequencies := [101]int{}
	for _, num := range nums {
		frequencies[num]++
	}
	sum := 0
	for i := 1; i <= 100; i++ {
		if frequencies[i]%k == 0 {
			sum += frequencies[i] * i
		}
	}
	return sum
}
