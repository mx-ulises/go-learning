package leetcode

func getProduct(n int) int {
	product := 1
	for 0 < n {
		product *= n % 10
		n /= 10
	}
	return product
}

func smallestNumber(n int, t int) int {
	product := getProduct(n)
	for product%t != 0 {
		n++
		product = getProduct(n)
	}
	return n
}
