package leetcode

func furthestDistanceFromOrigin(moves string) int {
	r, l, m := 0, 0, 0
	for _, move := range moves {
		if move == 'R' {
			r++
		} else if move == 'L' {
			l++
		} else if move == '_' {
			m++
		}
	}
	return max(r+m-l, l+m-r)
}
