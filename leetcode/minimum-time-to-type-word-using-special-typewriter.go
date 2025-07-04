package leetcode

func distance(a, b int) int {
	d := a - b
	if d < 0 {
		d = -d
	}
	return min(d, 26-d)
}

func minTimeToType(word string) int {
	t, prev := 0, 0
	for i := 0; i < len(word); i++ {
		current := int(word[i] - 'a')
		t += distance(prev, current) + 1
		prev = current
	}
	return t
}
