func InArray(nums *[]int, num int) bool {
	n := len(*nums)
	for i := 0; i < n; i++ {
		if (*nums)[i] == num {
			return true
		}
	}
	return false
}

func getSneakyNumbers(nums []int) []int {
	output := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] != i {
			j := nums[i]
			if nums[j] == j {
				if InArray(&output, j) == false {
					output = append(output, j)
				}
				break
			}
			nums[i] = nums[j]
			nums[j] = j
		}
	}
	return output
}
