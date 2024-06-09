func Adjust(num, k int) int {
	if num < 0 {
		num += (1 - (num / k)) * k
	}
	return num
}

func subarraysDivByK(nums []int, k int) int {
	divisorCount := map[int]int{0: 1}
	current := 0
	count := 0
	for _, num := range nums {
		current += Adjust(num, k)
		key := current % k
		count += divisorCount[key]
		divisorCount[key]++
	}
	return count
}
