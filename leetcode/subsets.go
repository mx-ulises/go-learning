func GetSubsets(output *[][]int, nums *[]int, preffix *[]int) {
	if len(*nums) == 0 {
		subset := make([]int, len(*preffix))
		copy(subset, (*preffix))
		(*output) = append(*output, subset)
	} else {
		num := (*nums)[len(*nums)-1]
		(*nums) = (*nums)[:len(*nums)-1]
		GetSubsets(output, nums, preffix)
		(*preffix) = append((*preffix), num)
		GetSubsets(output, nums, preffix)
		(*preffix) = (*preffix)[:len(*preffix)-1]
		(*nums) = append(*nums, num)
	}
}

func subsetsRecursive(nums []int) [][]int {
	output := make([][]int, 0)
	preffix := make([]int, 0)
	GetSubsets(&output, &nums, &preffix)
	return output
}

func subsetsIterative(nums []int) [][]int {
	output := make([][]int, 0)
	output = append(output, []int{})
	for i := 0; i < len(nums); i++ {
		currentElements := len(output)
		for j := 0; j < currentElements; j++ {
			subset := make([]int, len(output[j]))
			copy(subset, output[j])
			subset = append(subset, nums[i])
			output = append(output, subset)
		}
	}
	return output
}

func subsets(nums []int) [][]int {
	return subsetsIterative(nums)
}
