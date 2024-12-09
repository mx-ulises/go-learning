func isArraySpecial(nums []int, queries [][]int) []bool {
	n, m := len(nums), len(queries)
	breaches := make([]int, n)
	for i := 1; i < n; i++ {
		j := i - 1
		breach := (nums[j] ^ nums[i] ^ 1) & 1
		breaches[i] = breaches[j] + breach
	}
	output := make([]bool, m)
	for i, query := range queries {
		output[i] = breaches[query[0]] == breaches[query[1]]
	}
	return output
}
