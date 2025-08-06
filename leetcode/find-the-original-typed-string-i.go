package leetcode

func possibleStringCount(word string) int {
	total := 1
	current, currentCount := '0', 1
	for _, c := range word {
		if current == c {
			currentCount++
		} else {
			total += currentCount - 1
			currentCount = 1
			current = c
		}
	}
	total += currentCount - 1
	return total
}
