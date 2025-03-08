func minimumRecolors(blocks string, k int) int {
	l, r := 0, 0
	count := 0
	for r < k {
		if blocks[r] == 'W' {
			count++
		}
		r++
	}
	minimal := count
	for r < len(blocks) {
		if blocks[r] == 'W' {
			count++
		}
		if blocks[l] == 'W' {
			count--
		}
		minimal = min(minimal, count)
		r++
		l++
	}
	return minimal
}
