func GetMaximal(nums *[]int) int {
	maximal := (*nums)[0]
	for _, num := range *nums {
		maximal = max(maximal, num)
	}
	return maximal
}

func maximizeSum(nums []int, k int) int {
	return k*GetMaximal(&nums) + k*(k-1)/2
}
