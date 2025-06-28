package leetcode

func replaceElements(arr []int) []int {
	current := -1
	for i := len(arr) - 1; 0 <= i; i-- {
		next := max(current, arr[i])
		arr[i] = current
		current = next
	}
	return arr
}
