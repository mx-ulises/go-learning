func GetNum(x int) int {
	return x % 1_000_000
}

func findDuplicates(nums []int) []int {
	n := len(nums)
	output := []int{}
	for i := 0; i < n; i++ {
		num := GetNum(nums[i])
		j := num - 1
		nums[j] += 1_000_000
		if 2_000_000 < nums[j] {
			output = append(output, num)
		}
	}
	return output
}
