package leetcode

func maxProduct(n int) int {
	first, second, d := 0, 0, -1
	for 0 < n {
		d, n = n%10, n/10
		if first <= d {
			second, first = first, d
		} else if second < d {
			second = d
		}
	}
	return first * second
}
