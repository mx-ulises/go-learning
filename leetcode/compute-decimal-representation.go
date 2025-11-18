package leetcode

func decimalRepresentation(n int) []int {
	components := getComponents(n)
	return reverseComponents(components)
}

func getComponents(n int) []int {
	components := []int{}
	multiplier := 1
	for 0 < n {
		components = addComponent(components, (n%10)*multiplier)
		n, multiplier = n/10, multiplier*10
	}
	return components
}

func addComponent(components []int, component int) []int {
	if 0 < component {
		components = append(components, component)
	}
	return components
}

func reverseComponents(components []int) []int {
	n := len(components)
	for left := 0; left < n/2; left++ {
		right := n - left - 1
		components[left], components[right] = components[right], components[left]
	}
	return components
}
