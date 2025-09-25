package leetcode

func greatestLetter(s string) string {
	visited := [26]int{}
	maximal := byte('0')
	for i := 0; i < len(s); i++ {
		c, j := s[i], -1
		if 'a' <= c && c <= 'z' {
			j = int(c - 'a')
			visited[j] |= 1
		} else {
			j = int(c - 'A')
			visited[j] |= 2
		}
		if visited[j] == 3 {
			maximal = max(maximal, byte(j)+'A')
		}
	}
	if maximal == '0' {
		return ""
	}
	return string([]byte{maximal})
}
