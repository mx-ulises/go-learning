func sequentialDigits(low int, high int) []int {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8}
	output := []int{}
	for 0 < len(elements) {
		element := elements[0]
		elements = elements[1:]
		if high < element {
			continue
		}
		if low <= element {
			output = append(output, element)
		}
		d := element % 10
		if d != 9 {
			elements = append(elements, (element*10)+d+1)
		}
	}
	return output
}
