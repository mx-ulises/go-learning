func productExceptSelf(nums []int) []int {
	n := len(nums)
	products := make([]int, n)
	currentProduct := 1
	for i := 0; i < n; i++ {
		products[i] = currentProduct
		currentProduct *= nums[i]
	}
	currentProduct = 1
	for i := (n - 1); 0 <= i; i-- {
		products[i] *= currentProduct
		currentProduct *= nums[i]
	}
	return products
}
